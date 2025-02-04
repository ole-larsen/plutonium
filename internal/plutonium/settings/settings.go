package settings

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type OAUTH2 struct {
	Callback string `mapstructure:"OAUTH2_CALLBACK"`
	Provider string `mapstructure:"OAUTH2_PROVIDER"`
}

// Settings represents the configuration settings for the server.
type Settings struct {
	OAUTH2     OAUTH2
	MarketName string `mapstructure:"MARKETNAME"`
	Domain     string `mapstructure:"DOMAIN"`
	DSN        string
	PrivateKey string `mapstructure:"ETH_PRIVATE_KEY"`
	Network    string `mapstructure:"ETH_NETWORK"`
	XToken     string `mapstructure:"X_TOKEN"`
	Secret     string `mapstructure:"SECRET"`
	DB         DB
	GRPC       GRPC
}

type GRPC struct {
	Host string `mapstructure:"GRPC_HOST"`
	Cert string `mapstructure:"GRPC_CERT"`
	Key  string `mapstructure:"GRPC_KEY"`
	Port int    `mapstructure:"GRPC_PORT"`
}

type DB struct {
	Proto    string `mapstructure:"DB_SQL_PROTO"`
	Host     string `mapstructure:"DB_SQL_HOST"`
	Port     string `mapstructure:"DB_SQL_PORT"`
	Username string `mapstructure:"DB_SQL_USERNAME"`
	Password string `mapstructure:"DB_SQL_PASSWORD"`
	Database string `mapstructure:"DB_SQL_DATABASE"`
	SSL      bool   `mapstructure:"DB_SQL_SSL"`
}

var (
	config = &Settings{}
	once   sync.Once
)

// LoadConfig initializes and returns the settings singleton.
func LoadConfig(cfgPath string) *Settings {
	once.Do(func() {
		config = InitConfig(cfgPath,
			WithMarketname(),
			WithDomain(),
			WithDSN(),
			WithEthPK(),
			WithEthWeb(),
			WithXToken(),
			WithSecret(),
			WithOauth2Provider(),
			WithOauth2Callback(),
			WithGRPC(),
			WithGRPCCert(),
			WithGRPCKey(),
		)
	})

	return config
}

// InitConfig initializes the settings from the specified config file.
func InitConfig(cfgPath string, opts ...func(*Settings)) *Settings {
	if err := LoadEnv(cfgPath); err != nil {
		panic(err)
	}

	config.Reload(opts...)

	return config
}

// LoadEnv sets the configuration file path and reads the configuration.
// It initializes the viper configuration with the specified file.
func LoadEnv(cfgPath string) error {
	viper.SetConfigFile(cfgPath)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error while reading config file: %w", err)
	}

	return nil
}

// Reload applies configuration options to the settings.
func (c *Settings) Reload(opts ...func(*Settings)) {
	for _, opt := range opts {
		opt(c)
	}
}

func WithMarketname() func(*Settings) {
	return func(ss *Settings) {
		name := viper.GetString("MARKETNAME")
		ss.MarketName = name
	}
}

func WithDomain() func(*Settings) {
	return func(ss *Settings) {
		name := viper.GetString("DOMAIN")
		ss.MarketName = name
	}
}

func WithOauth2Provider() func(*Settings) {
	return func(ss *Settings) {
		provider := viper.GetString("OAUTH2_PROVIDER")
		ss.OAUTH2.Provider = provider
	}
}

func WithOauth2Callback() func(*Settings) {
	return func(ss *Settings) {
		callback := viper.GetString("OAUTH2_CALLBACK")
		ss.OAUTH2.Callback = callback
	}
}

// WithDSN is a functional option for building the DSN string.
func WithDSN() func(*Settings) {
	return func(ss *Settings) {
		ss.DB = DB{
			Proto:    viper.GetString("DB_SQL_PROTO"),
			Host:     viper.GetString("DB_SQL_HOST"),
			Port:     viper.GetString("DB_SQL_PORT"),
			Username: viper.GetString("DB_SQL_USERNAME"),
			Password: viper.GetString("DB_SQL_PASSWORD"),
			Database: viper.GetString("DB_SQL_DATABASE"),
		}

		dsn := fmt.Sprintf(
			"%s://%s:%s@%s:%s/%s",
			ss.DB.Proto, ss.DB.Username, ss.DB.Password, ss.DB.Host, ss.DB.Port, ss.DB.Database,
		)
		if !viper.GetBool("DB_SQL_SSL") {
			dsn += "?sslmode=disable"
		}

		ss.DSN = dsn
	}
}

// WithEthPK is a functional option for building the pk string.
func WithEthPK() func(*Settings) {
	return func(ss *Settings) {
		pk := viper.GetString("ETH_PRIVATE_KEY")
		ss.PrivateKey = pk
	}
}

// WithEthWeb is a functional option for building the pk string.
func WithEthWeb() func(*Settings) {
	return func(ss *Settings) {
		network := viper.GetString("ETH_NETWORK")
		ss.Network = network
	}
}

// WithXToken is a functional option for initializing the XToken field in the Settings struct.//+.
func WithXToken() func(*Settings) {
	return func(ss *Settings) {
		token := viper.GetString("XTOKEN")
		ss.XToken = token
	}
}

func WithSecret() func(*Settings) {
	return func(ss *Settings) {
		token := viper.GetString("SECRET")
		ss.Secret = token
	}
}

func WithGRPC() func(*Settings) {
	return func(ss *Settings) {
		addr := strings.Split(viper.GetString("GRPC_URL"), ":")

		const reqLen = 2

		if len(addr) != reqLen {
			panic(fmt.Errorf("wrong a parameters"))
		}

		ss.GRPC.Host = addr[0]

		port, err := strconv.Atoi(addr[1])
		if err != nil {
			panic(fmt.Errorf("wrong a parameters"))
		}

		ss.GRPC.Port = port
	}
}

func WithGRPCCert() func(*Settings) {
	return func(ss *Settings) {
		crt := viper.GetString("GRPC_CERT")
		ss.GRPC.Cert = crt
	}
}

func WithGRPCKey() func(*Settings) {
	return func(ss *Settings) {
		key := viper.GetString("GRPC_KEY")
		ss.GRPC.Key = key
	}
}
