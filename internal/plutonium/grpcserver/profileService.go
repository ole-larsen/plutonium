package grpcserver

import (
	"context"

	"connectrpc.com/connect"
	profilev1 "github.com/ole-larsen/plutonium/gen/profile/v1"
	"github.com/ole-larsen/plutonium/gen/profile/v1/profilev1connect"
	"github.com/ole-larsen/plutonium/internal/blockchain"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium/httpclient"
	"github.com/ole-larsen/plutonium/internal/plutonium/oauth2client"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/ole-larsen/plutonium/internal/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProfileServiceServer implements the MarketService API.
type ProfileServiceServer struct {
	profilev1connect.UnimplementedProfileServiceHandler
	logger     *log.Logger
	storage    storage.DBStorageInterface
	web3Dialer *blockchain.Web3Dialer
	httpDialer *httpclient.HTTPClient
	settings   *settings.Settings
	oauth2     *oauth2client.Oauth2
}

func (s *ProfileServiceServer) SetSettings(cfg *settings.Settings) *ProfileServiceServer {
	s.settings = cfg
	return s
}

func (s *ProfileServiceServer) SetLogger(logger *log.Logger) *ProfileServiceServer {
	s.logger = logger
	return s
}

func (s *ProfileServiceServer) SetStorage(store storage.DBStorageInterface) *ProfileServiceServer {
	s.storage = store
	return s
}

func (s *ProfileServiceServer) SetWeb3Dialer(web3Dialer *blockchain.Web3Dialer) *ProfileServiceServer {
	s.web3Dialer = web3Dialer
	return s
}

func (s *ProfileServiceServer) SetHTTPDialer(httpDialer *httpclient.HTTPClient) *ProfileServiceServer {
	s.httpDialer = httpDialer
	return s
}

func (s *ProfileServiceServer) SetOauth2(oauth2cfg *oauth2client.Oauth2) *ProfileServiceServer {
	s.oauth2 = oauth2cfg
	return s
}

func (s *ProfileServiceServer) PatchUser(
	ctx context.Context,
	request *connect.Request[profilev1.PatchUserRequest],
) (*connect.Response[profilev1.PatchUserResponse], error) {
	body := request.Msg.Body

	exists, err := s.storage.GetUsersRepository().GetUserByID(ctx, body.Id)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}

	if exists == nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.NotFound, "user not found")

	}
	/*
		Csrf          string                 `protobuf:"bytes,1,opt,name=csrf,proto3" json:"csrf,omitempty"`
		Id            int64                  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
		Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
		Address       string                 `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
		Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
		Gravatar      string
	*/
	if body.Csrf == "" {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.InvalidArgument, "csrf token is required")
	}
	userMap := make(map[string]interface{})

	if body.Id != 0 {
		userMap["id"] = body.Id
	}
	if body.Username != "" {
		userMap["username"] = body.Username
	}
	if body.Email != "" {
		userMap["email"] = body.Email
	}
	if body.Gravatar != "" {
		userMap["gravatar"] = body.Gravatar
	}
	user, err := s.storage.GetUsersRepository().Update(ctx, userMap)
	if err != nil {
		s.logger.Errorln(NewError(err))
		return nil, status.Error(codes.Internal, err.Error())
	}
	publicUser := NestPublicUser(user)
	response := &profilev1.PatchUserResponse{
		Response: &profilev1.PatchUserResponse_User{
			User: publicUser,
		},
	}
	return connect.NewResponse(response), status.Errorf(codes.OK, "OK")
}
