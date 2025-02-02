package grpcserver

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"connectrpc.com/connect"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	marketv1 "github.com/ole-larsen/plutonium/gen/market/v1"
	"github.com/ole-larsen/plutonium/gen/market/v1/marketv1connect"
	"github.com/ole-larsen/plutonium/internal/blockchain"
	"github.com/ole-larsen/plutonium/internal/hash"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium/httpclient"
	"github.com/ole-larsen/plutonium/internal/plutonium/oauth2client"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/internal/storage"
	"github.com/ole-larsen/plutonium/internal/storage/db/repository"
	"github.com/ole-larsen/plutonium/models"
	"github.com/sethvargo/go-password/password"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

const pwdLength = 8
const numDigits = 4
const numSymbols = 4
const userNotFoundMsg = "[repository]: user not found"
const defaultGravatarSize = 200

// MarketServiceServer implements the MarketService API.
type MarketServiceServer struct {
	marketv1connect.UnimplementedMarketServiceHandler
	logger     *log.Logger
	storage    storage.DBStorageInterface
	web3Dialer *blockchain.Web3Dialer
	httpDialer *httpclient.HTTPClient
	settings   *settings.Settings
	oauth2     *oauth2client.Oauth2
}

func (s *MarketServiceServer) SetSettings(cfg *settings.Settings) *MarketServiceServer {
	s.settings = cfg
	return s
}

func (s *MarketServiceServer) SetLogger(logger *log.Logger) *MarketServiceServer {
	s.logger = logger
	return s
}

func (s *MarketServiceServer) SetStorage(store storage.DBStorageInterface) *MarketServiceServer {
	s.storage = store
	return s
}

func (s *MarketServiceServer) SetWeb3Dialer(web3Dialer *blockchain.Web3Dialer) *MarketServiceServer {
	s.web3Dialer = web3Dialer
	return s
}

func (s *MarketServiceServer) SetHTTPDialer(httpDialer *httpclient.HTTPClient) *MarketServiceServer {
	s.httpDialer = httpDialer
	return s
}

func (s *MarketServiceServer) SetOauth2(oauth2cfg *oauth2client.Oauth2) *MarketServiceServer {
	s.oauth2 = oauth2cfg
	return s
}

func (s *MarketServiceServer) Contracts(
	ctx context.Context,
	_ *connect.Request[marketv1.ContractsRequest],
) (*connect.Response[marketv1.ContractsResponse], error) {
	success := &marketv1.Success{}

	if err := s.web3Dialer.Load(ctx); err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	marketPlace := &marketv1.PublicMarketplaceContract{
		Name:    s.web3Dialer.Market.Marketplace.Name,
		Abi:     s.web3Dialer.Market.Marketplace.ABI,
		Address: s.web3Dialer.Market.Marketplace.Address.String(),
		Fee:     wrapperspb.String(s.web3Dialer.Market.Marketplace.Fee.String()),
		Owner:   s.web3Dialer.Market.Marketplace.Owner.String(),
	}

	success.Contracts = &marketv1.PublicContracts{
		Marketplace: marketPlace,
	}

	if len(s.web3Dialer.Market.Collections) > 0 {
		collections := make(map[string]*marketv1.PublicContract)
		for collectionID, collection := range s.web3Dialer.Market.Collections {
			collections[collectionID] = &marketv1.PublicContract{
				Name:    collection.Name,
				Abi:     collection.ABI,
				Address: collection.Address.String(),
			}
		}

		success.Contracts.Collections = collections
	}

	if len(s.web3Dialer.Market.Auctions) > 0 {
		auctions := make(map[string]*marketv1.PublicContract)
		for _, auction := range s.web3Dialer.Market.Auctions {
			auctions[auction.Name] = &marketv1.PublicContract{
				Name:    auction.Name,
				Abi:     auction.ABI,
				Address: auction.Address.String(),
			}
		}

		success.Contracts.Auctions = auctions
	}

	response := &marketv1.ContractsResponse{
		Response: &marketv1.ContractsResponse_Data{
			Data: success,
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *MarketServiceServer) Categories(
	ctx context.Context,
	_ *connect.Request[marketv1.CategoriesRequest],
) (*connect.Response[marketv1.CategoriesResponse], error) {
	items, err := s.storage.
		GetCategoriesRepository().
		GetPublicCollectibleCategories(ctx, s.storage.GetUsersRepository())
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &marketv1.CategoriesResponse{
		Response: &marketv1.CategoriesResponse_Data{
			Data: &marketv1.SuccessCategories{
				Categories: NestPublicCategories(items),
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *MarketServiceServer) Nonce(
	ctx context.Context,
	request *connect.Request[marketv1.NonceRequest],
) (*connect.Response[marketv1.NonceResponse], error) {
	address := strings.ToLower(request.Msg.Address)

	if address == "" {
		err := NewError(fmt.Errorf("address is required"))
		return nil, status.Error(codes.Internal, err.Error())
	}

	user, err := s.storage.GetUsersRepository().GetUserByAddress(ctx, address)

	if err != nil && err.Error() != "[repository]: user not found" {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	// to get payload from current user:
	// 1. check user exist in database
	// 2. if request causes error, return auth error
	// 3. if user exist, return existing payload
	// 4. If user is not exists, register new one
	if user == nil {
		if err = Register(ctx, address, s.storage.GetUsersRepository(), s.settings.Secret); err != nil {
			return nil, status.Error(codes.NotFound, err.Error())
		}
	}

	// fetch new payload
	payload, err := HandleNonce(ctx, address, s.storage.GetUsersRepository())
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	response := &marketv1.NonceResponse{
		Response: &marketv1.NonceResponse_Data{
			Data: &marketv1.Nonce{
				Address: address,
				Nonce:   payload.Nonce,
				Uuid:    payload.UUID,
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}

func (s *MarketServiceServer) Verify(
	ctx context.Context,
	request *connect.Request[marketv1.VerifyRequest],
) (*connect.Response[marketv1.VerifyResponse], error) {
	address := strings.ToLower(request.Msg.Address)

	if address == "" {
		err := NewError(fmt.Errorf("address is required"))
		return nil, status.Error(codes.Internal, err.Error())
	}

	user, err := s.storage.GetUsersRepository().GetUserByAddress(ctx, address)

	if err != nil && err.Error() != "[repository]: user not found" {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		if err.Error() == userNotFoundMsg {
			s.logger.Errorln(NewError(err))
			return nil, status.Error(codes.NotFound, err.Error())
		}

		s.logger.Errorln(NewError(err))

		return nil, status.Error(codes.Internal, err.Error())
	}

	nonce := ""
	if user.Nonce.Valid {
		nonce = user.Nonce.String
	}

	userUUID := ""
	if user.UUID.Valid {
		userUUID = user.UUID.String
	}

	marketName := s.settings.MarketName

	msg, err := hexutil.Decode(request.Msg.Msg)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if strings.Contains(string(msg), marketName) &&
		strings.Contains(string(msg), address) &&
		strings.Contains(string(msg), nonce) {
		userMap := make(map[string]interface{})
		userMap["id"] = user.ID
		userMap["nonce"] = generateNonce()

		if err = s.storage.GetUsersRepository().UpdateNonce(ctx, userMap); err != nil {
			s.logger.Errorln(NewError(err))
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	signature := request.Msg.Signature

	token, err := s.GetAccessToken(ctx, user.Email, signature)

	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	var wallpaper *models.PublicFile
	if user.Wallpaper != nil {
		wallpaper = &models.PublicFile{
			Attributes: user.Wallpaper.Attributes,
			ID:         user.Wallpaper.ID,
		}
	}

	response := &marketv1.VerifyResponse{
		Response: &marketv1.VerifyResponse_Data{
			Data: &marketv1.VerifiedAccess{
				User: &marketv1.PublicUser{
					Id: user.ID,
					Attributes: &marketv1.PublicUserAttributes{
						Address:   address,
						Uuid:      userUUID,
						Email:     user.Email,
						Gravatar:  user.Gravatar,
						Wallpaper: NestMarketPublicFile(wallpaper),
						Username:  user.Username,
					},
				},
				Token: &marketv1.Oauth2Token{
					AccessToken:  token.AccessToken,
					Code:         token.Code,
					Expiry:       timestamppb.New(token.Expiry),
					OriginalUrl:  token.OriginalURL,
					RefreshToken: token.RefreshToken,
					State:        token.State,
					TokenType:    token.TokenType,
				},
			},
		},
	}

	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}
func Register(ctx context.Context, address string, users *repository.UsersRepository, secret string) error {
	pwd, err := password.Generate(pwdLength, numDigits, numSymbols, false, false)
	if err != nil {
		return err
	}

	hashPwd, err := hash.PasswordWithSecret(pwd, secret)
	if err != nil {
		return err
	}

	email := address + "@" + "plutonium"
	userMap := make(map[string]interface{})
	userMap["username"] = address
	userMap["email"] = email
	userMap["password"] = hashPwd
	userMap["secret"] = "" // JNUGNHA27JMIHA5I
	userMap["address"] = "{" + address + "}"
	userMap["nonce"] = generateNonce()
	userMap["gravatar"] = gravatar(email, defaultGravatarSize)
	// generate secret per user
	length := 16
	userMap["rsa_secret"] = hash.RandStringBytes(length)

	return users.Create(ctx, userMap)
}

func HandleNonce(ctx context.Context, address string, repo *repository.UsersRepository) (*models.Nonce, error) {
	user, err := repo.GetUserByAddress(ctx, address)

	if err != nil && err.Error() != userNotFoundMsg {
		return nil, err
	}

	payload := models.Nonce{
		Address: address,
	}
	if user.Nonce.Valid {
		payload.Nonce = user.Nonce.String
	}

	if user.UUID.Valid {
		payload.UUID = user.UUID.String
	}

	return &payload, nil
}

func generateNonce() string {
	id := uuid.New()
	return id.String()
}

func gravatar(email string, size int) string {
	gravatarURL := "https://gravatar.com/avatar/"
	if email != "" {
		return gravatarURL + hash.GetMD5Hash(email) + "?s=" + strconv.Itoa(size) + "&d=retro"
	}

	return gravatarURL + "?s=" + strconv.Itoa(size) + "&d=retro"
}

func (s *MarketServiceServer) GetAccessToken(ctx context.Context, email, signature string) (*models.Callback, error) {
	s.logger.Infoln("here is signature", signature)

	credentials, err := s.httpDialer.GetCredentials(ctx, email)
	if err != nil {
		s.logger.Errorln(err)
		return nil, err
	}

	clientID := *credentials.ClientID
	clientSecret := *credentials.ClientSecret

	config := s.oauth2.Client.Config(clientID, clientSecret)

	s.oauth2.Config[clientID] = config

	authURL := s.oauth2.Client.AuthorizeURL(&config)

	return s.httpDialer.Authorize(ctx, authURL)
}
