package plutonium

import (
	"github.com/ole-larsen/plutonium/internal/blockchain"
	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/storage"

	"github.com/ole-larsen/plutonium/internal/plutonium/grpcserver"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
)

// Server represents the main server structure.
type Server struct {
	settings   *settings.Settings
	logger     *log.Logger
	storage    storage.DBStorageInterface
	grpc       *grpcserver.GRPCServer
	web3Dialer *blockchain.Web3Dialer
}

// NewServer creates and returns a new Server instance.
// Initially, the Server instance is created with nil settings.
func NewServer() *Server {
	return &Server{}
}

// SetSettings configures the server with the given settings.
//
// It accepts a pointer to a settings.Config object, assigns it to
// the server, and returns the updated Server instance.
func (s *Server) SetSettings(cfg *settings.Settings) *Server {
	s.settings = cfg
	return s
}

func (s *Server) SetLogger(logger *log.Logger) *Server {
	s.logger = logger
	return s
}

// SetStorage sets the storage interface used by the server.
func (s *Server) SetStorage(store storage.DBStorageInterface) *Server {
	s.storage = store
	return s
}

// SetGRPC sets the gRPC service used by the server.
func (s *Server) SetGRPC(grpc *grpcserver.GRPCServer) *Server {
	s.grpc = grpc
	return s
}

func (s *Server) SetWeb3Dialer(web3Dialer *blockchain.Web3Dialer) {
	s.web3Dialer = web3Dialer
}

// GetSettings retrieves the current settings configuration of the server.
func (s *Server) GetSettings() *settings.Settings {
	return s.settings
}

// GetLogger retrieves the logger associated with the server.
func (s *Server) GetLogger() *log.Logger {
	return s.logger
}

// GetStorage retrieves the storage used by the server.
func (s *Server) GetStorage() storage.DBStorageInterface {
	return s.storage
}

// GetGRPC retrieves the gRPC service used by the server.
func (s *Server) GetGRPC() *grpcserver.GRPCServer {
	return s.grpc
}

func (s *Server) GetWeb3Dialer() *blockchain.Web3Dialer {
	return s.web3Dialer
}
