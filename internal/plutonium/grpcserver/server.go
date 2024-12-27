// Package grpcserver provides the implementation of a gRPC server with support
// for both HTTP/1.1 and HTTP/2. It handles user authentication, file transfer,
// and key-value storage operations through the corresponding gRPC services.
//
// This package offers both standard and TLS-encrypted server configurations,
// customizable through the provided configuration structure. It supports
// CORS and includes JWT-based authentication interceptors for enhanced security.
package grpcserver

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/ole-larsen/plutonium/internal/log"
	"github.com/ole-larsen/plutonium/internal/plutonium/settings"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const timeout = 3

/*
Overview:
The `grpcserver` package defines the `GRPCServer` struct, which is responsible for
setting up and starting a gRPC server that supports multiple services, including
user authentication, key-value management, and file transfer. The package also provides
configuration utilities for setting up TLS, managing CORS, and customizing the server
host, port, and protocol.

Key Components:
- GRPCServer: The main server struct that encapsulates configuration, logging,
  and storage.
- Interface: A generic interface that allows for customization of the server's
  host, port, protocol, and storage.
- TLS Support: Functions for loading TLS certificates and enabling encrypted
  communication.
- CORS Support: CORS configuration to handle cross-origin requests.
- Services: The package integrates three main services—user, key-value (KV), and file
  transfer—each accessible through gRPC endpoints.

*/

// GRPC defines the contract for setting up and running a gRPC server.
type GRPC interface {
	// SetHost sets the server's hostname.
	SetHost(h string) *GRPCServer
	// SetPort sets the server's port number.
	SetPort(p int) *GRPCServer
	// GetHost retrieves the configured host.
	GetHost() string
	// GetPort retrieves the configured port.
	GetPort() int
	// ListenAndServe starts the server without TLS.
	ListenAndServe(cfg *settings.Settings) error
	// ListenAndServeTLS starts the server with TLS.
	ListenAndServeTLS(cfg *settings.Settings) error
	// LoadTLS loads TLS certificates and returns a TLS configuration.
	LoadTLS(cfg *settings.Settings) (*tls.Config, error)
	// ListenTLS creates a TLS-enabled HTTP server.
	ListenTLS(cfg *settings.Settings) (*http.Server, error)
}

// GRPCServer implements the Interface and is responsible for managing the gRPC services.
type GRPCServer struct {
	logger *log.Logger
	host   string
	port   int
}

// NewGRPCServer creates a new instance of GRPCServer with default logger settings.
func NewGRPCServer() *GRPCServer {
	return &GRPCServer{
		logger: log.NewLogger("info", log.DefaultBuildLogger),
	}
}

// SetHost sets the host for the server.
func (s *GRPCServer) SetHost(h string) *GRPCServer {
	s.host = h
	return s
}

// SetPort sets the port for the server.
func (s *GRPCServer) SetPort(p int) *GRPCServer {
	s.port = p
	return s
}

// GetHost returns the host configured for the server.
func (s *GRPCServer) GetHost() string {
	return s.host
}

// GetPort returns the port configured for the server.
func (s *GRPCServer) GetPort() int {
	return s.port
}

// ListenAndServeTLS starts the server with TLS encryption.
func (s *GRPCServer) ListenAndServeTLS(cfg *settings.Settings) error {
	srv, err := s.ListenTLS(cfg)

	if err != nil {
		return fmt.Errorf("failed to start server gRPC: %w", err)
	}

	fmt.Printf("server is listening on https://%s:%d\n", cfg.GRPC.Host, cfg.GRPC.Port)

	return srv.ListenAndServeTLS("", "")
}

// ListenAndServe starts the server without TLS encryption.
func (s *GRPCServer) ListenAndServe(cfg *settings.Settings) error {
	srv, err := s.Listen(cfg)

	if err != nil {
		return fmt.Errorf("failed to start server gRPC: %w", err)
	}

	fmt.Printf("server is listening on http://%s:%d\n", cfg.GRPC.Host, cfg.GRPC.Port)

	return srv.ListenAndServe()
}

// LoadTLS loads and returns a TLS configuration using server certificates.
func (s *GRPCServer) LoadTLS(cfg *settings.Settings) (*tls.Config, error) {
	// Load the TLS certificate and key from the bytes
	cert, err := tls.X509KeyPair([]byte(cfg.GRPC.Cert), []byte(cfg.GRPC.Key))
	if err != nil {
		return nil, fmt.Errorf("failed to load certificate: %w", err)
	}

	// Create a TLS Config with the loaded certificate
	return &tls.Config{
		MinVersion:   tls.VersionTLS13, // Use TLS 1.3 or TLS 1.2 as the minimum version
		Certificates: []tls.Certificate{cert},
		// Optional: Ensure that only HTTP/2 connections are allowed
		NextProtos: []string{"h2"},
	}, nil
}

// Listen configures and returns an HTTP/2-enabled server.
func (s *GRPCServer) Listen(_ *settings.Settings) (*http.Server, error) {
	if s == nil || s.host == "" || s.port == 0 {
		return nil, NewError(fmt.Errorf("server gRPC is disabled"))
	}

	address := fmt.Sprintf("%s:%d", s.host, s.port)

	mux := http.NewServeMux()

	return &http.Server{
		Addr:              address,
		ReadHeaderTimeout: timeout * time.Second,
		Handler:           h2c.NewHandler(NewCORS().Handler(mux), &http2.Server{}),
	}, nil
}

// ListenTLS configures and returns a TLS-enabled server.
func (s *GRPCServer) ListenTLS(cfg *settings.Settings) (*http.Server, error) {
	srv, err := s.Listen(cfg)
	if err != nil {
		return nil, err
	}

	tlsConfig, err := s.LoadTLS(cfg)
	if err != nil {
		return nil, err
	}

	srv.TLSConfig = tlsConfig

	return srv, nil
}

// SetupGRPC is a helper function to configure a new GRPCServer.
func SetupGRPC(host string, port int) *GRPCServer {
	if host == "" || port == 0 {
		return nil
	}

	return NewGRPCServer().
		SetHost(host).
		SetPort(port)
}

// NewCORS configures CORS (Cross-Origin Resource Sharing) for the server.
func NewCORS() *cors.Cors {
	// setup cors for access from another sources
	return cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowOriginFunc: func(_ /* origin */ string) bool {
			// Allow all origins, which effectively disables CORS.
			return true
		},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{
			// Content-Type is in the default safelist.
			"Accept",
			"Accept-Encoding",
			"Accept-Post",
			"Connect-Accept-Encoding",
			"Connect-Content-Encoding",
			"Content-Encoding",
			"Grpc-Accept-Encoding",
			"Grpc-Encoding",
			"Grpc-Message",
			"Grpc-Status",
			"Grpc-Status-Details-Bin",
		},
		// Let browsers cache CORS information for longer, which reduces the number
		// of preflight requests. Any changes to ExposedHeaders won't take effect
		// until the cached data expires. FF caps this value at 24h, and modern
		// Chrome caps it at 2h.
		MaxAge: int(2 * time.Hour / time.Second),
	})
}
