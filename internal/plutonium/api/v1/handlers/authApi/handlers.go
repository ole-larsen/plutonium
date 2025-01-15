package authapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/go-openapi/runtime/middleware"
	"github.com/ole-larsen/plutonium/internal/hash"
	"github.com/ole-larsen/plutonium/internal/plutonium"
	"github.com/ole-larsen/plutonium/internal/plutonium/jwt"
	"github.com/ole-larsen/plutonium/models"
	"github.com/ole-larsen/plutonium/restapi/operations/auth"
	"github.com/sethvargo/go-password/password"
	"golang.org/x/oauth2"
)

const UploadDir = "./uploads"

type API struct {
	service *plutonium.Server
}

type AuthAPI interface {
	GetWalletConnect(params auth.GetFrontendAuthWalletConnectParams, principal *models.Principal) middleware.Responder
	PostWalletConnect(params auth.PostFrontendAuthWalletConnectParams) middleware.Responder
	GetOauth2Callback(params auth.GetFrontendAuthCallbackParams) middleware.Responder
	Register(ctx context.Context, address string) error
	HandleNonce(ctx context.Context, address string) (*models.Nonce, error)
	GetAccessToken(email string) (string, error)
}

func NewAuthAPI(s *plutonium.Server) AuthAPI {
	return &API{service: s}
}

func (a *API) GetWalletConnect(params auth.GetFrontendAuthWalletConnectParams, principal *models.Principal) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if params.Operation == nil {
		return auth.NewGetFrontendAuthWalletConnectBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "operation is required",
		})
	}

	if params.Address == nil {
		return auth.NewGetFrontendAuthWalletConnectBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "address is required",
		})
	}

	address := strings.ToLower(*params.Address)

	operation := *params.Operation

	user, err := a.service.GetStorage().GetUsersRepository().GetUserByAddress(ctx, address)

	if err != nil && err.Error() != "[repository]: user not found" {
		return auth.NewGetFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if operation == "nonce" {
		// to get payload from current user:
		// 1. check user exist in database
		// 2. if request causes error, return auth error
		// 3. if user exist, return existing payload
		// 4. If user is not exists, register new one
		if user == nil {
			if err = a.Register(ctx, address); err != nil {
				return auth.NewGetFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
		}
	}

	// fetch new payload
	payload, err := a.HandleNonce(ctx, address)

	if err != nil {
		return auth.NewGetFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return auth.NewGetFrontendAuthWalletConnectOK().WithPayload(payload)
}

func (a *API) PostWalletConnect(params auth.PostFrontendAuthWalletConnectParams) middleware.Responder {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if params.Body == nil {
		return auth.NewPostFrontendAuthWalletConnectBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "body is required",
		})
	}

	user, err := a.service.GetStorage().GetUsersRepository().GetUserByAddress(ctx, params.Body.Address)

	if err != nil {
		if err.Error() == "[repository]: user not found" {
			return auth.NewGetFrontendAuthWalletConnectNotFound().WithPayload(&models.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
			})
		}
		return auth.NewPostFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	marketName := a.service.GetSettings().MarketName
	nonce := ""
	if user.Nonce.Valid {
		nonce = user.Nonce.String
	}
	uuid := ""
	if user.UUID.Valid {
		uuid = user.UUID.String
	}
	address := ""
	if len(user.Address) > 0 {
		address = user.Address[0]
	}

	msg, err := hexutil.Decode(params.Body.Msg)
	if err != nil {
		a.service.GetLogger().Errorln(err)
		return auth.NewPostFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if strings.Contains(string(msg), marketName) &&
		strings.Contains(string(msg), address) &&
		strings.Contains(string(msg), nonce) {

		userMap := make(map[string]interface{})
		userMap["id"] = user.ID
		userMap["nonce"] = generateNonce()

		if err = a.service.GetStorage().GetUsersRepository().UpdateNonce(ctx, userMap); err != nil {
			a.service.GetLogger().Errorln(err)
			return auth.NewPostFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		token, err := a.GetAccessToken(user.Email)

		if err != nil {
			a.service.GetLogger().Errorln(err)
			return auth.NewPostFrontendAuthWalletConnectInternalServerError().WithPayload(&models.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		payload := models.LoginMetamaskOK{
			Address:  address,
			UUID:     uuid,
			Email:    user.Email,
			ID:       user.ID,
			Nonce:    nonce,
			Token:    token,
			Username: user.Username,
		}

		return auth.NewPostFrontendAuthWalletConnectOK().WithPayload(&payload)
	}

	return auth.NewPostFrontendAuthWalletConnectUnauthorized().WithPayload(&models.ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: "unauthorized",
	})
}

func (a *API) Register(ctx context.Context, address string) error {
	pwd, err := password.Generate(8, 4, 4, false, false)
	if err != nil {
		return err
	}
	hashPwd, err := hashPassword(pwd, a.service.GetSettings().Secret)
	if err != nil {
		return err
	}

	userMap := make(map[string]interface{})
	userMap["username"] = address
	userMap["email"] = address + "@" + "plutonium"
	userMap["password"] = hashPwd
	userMap["secret"] = "" // JNUGNHA27JMIHA5I
	userMap["address"] = "{" + address + "}"
	userMap["nonce"] = generateNonce()

	// generate secret per user
	length := 16
	userMap["rsa_secret"] = hash.RandStringBytes(length)

	return a.service.GetStorage().GetUsersRepository().Create(ctx, userMap)
}

func (a *API) HandleNonce(ctx context.Context, address string) (*models.Nonce, error) {
	user, err := a.service.GetStorage().GetUsersRepository().GetUserByAddress(ctx, address)

	if err != nil && err.Error() != "[repository]: user not found" {
		return nil, err
	}

	if user != nil {
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

	return nil, nil
}

func (a *API) GetAccessToken(email string) (string, error) {
	credentials, err := a.service.GetHTTPDialer().GetCredentials(email)
	if err != nil {
		a.service.GetLogger().Errorln(err)
		return "", err
	}

	clientID := *credentials.ClientID
	clientSecret := *credentials.ClientSecret

	config := a.service.GetOauth2().Client.Config(clientID, clientSecret)

	a.service.GetOauth2().Config[clientID] = config

	authUrl := a.service.GetOauth2().Client.AuthorizeUrl(config)
	fmt.Println("Auth URL: ", authUrl)
	tkn, err := a.service.GetHTTPDialer().Authorize(authUrl)
	if err != nil {
		a.service.GetLogger().Errorln(err)
		return "", err
	}
	return jwt.CreateToken(email, tkn.AccessToken)
}

func (a *API) GetOauth2Callback(params auth.GetFrontendAuthCallbackParams) middleware.Responder {
	r := params.HTTPRequest

	clientID, err := a.service.GetOauth2().Client.GetClientIDFromReferer(r.Referer())
	if err != nil {
		return auth.NewPostFrontendAuthWalletConnectUnauthorized().WithPayload(&models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		})
	}

	config := a.service.GetOauth2().Config[clientID]

	err = r.ParseForm()

	if err != nil {
		a.service.GetLogger().Errorln(err)
		return auth.NewGetFrontendAuthCallbackInternalServerError().WithPayload(&models.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Message: "code is required",
		})
	}

	state := r.Form.Get("state")
	if state != "xyz" {
		return auth.NewGetFrontendAuthCallbackBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "state is required",
		})
	}

	code := r.Form.Get("code")
	if code == "" {
		return auth.NewGetFrontendAuthCallbackBadRequest().WithPayload(&models.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "code is required",
		})
	}

	token, err := config.Exchange(context.Background(), code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
	if err != nil {
		a.service.GetLogger().Errorln(err)
		return auth.NewGetFrontendAuthCallbackUnauthorized().WithPayload(&models.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Message: "unauthorized",
		})
	}

	payload := models.Callback{
		AccessToken:  token.AccessToken,
		Expiry:       token.Expiry,
		RefreshToken: token.RefreshToken,
		TokenType:    token.TokenType,
	}

	return auth.NewGetFrontendAuthCallbackOK().WithPayload(&payload)
}
