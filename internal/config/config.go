package config

import (
	"github.com/jinglanghe/go-start/utils/log"
	"github.com/spf13/viper"
)

var (
	Config AppConfig
)

type AppConfig struct {
	ApiPrefix  ApiPrefix
	RPCClient  RPCClient
	RPCServer  RPCServer
	HTTPServer HTTPServer
	HttpClient HttpClient
	Datasource string
	JWTConfig  JWTConfig
	Database   DbStruct
	Redis      RedisConfig
}

type ApiPrefix struct {
	Group string
}

// RPCClient is RPC client config.
type RPCClient struct {
	Host    map[string]string
	Dial    int
	Timeout int
}

// RPCServer is RPC server config.
type RPCServer struct {
	Network           string
	Addr              string
	Timeout           int
	IdleTimeout       int
	MaxLifeTime       int
	ForceCloseWait    int
	KeepAliveInterval int
	KeepAliveTimeout  int
}

// HTTPServer is http server config.
type HTTPServer struct {
	Network      string
	Addr         string
	ReadTimeout  int
	WriteTimeout int
}

// HttpClient is http client config.
type HttpClient struct {
	MaxIdleConn        int
	MaxConnPerHost     int
	MaxIdleConnPerHost int
	TimeoutSeconds     int
}

type JWTConfig struct {
	SignAlgorithm string
	SecretKey     string
	PublicKeyFile string
}

type DbStruct struct {
	ServerType  string
	Username    string
	Password    string
	Host        string
	Port        int
	DbName      string
	SslMode     string
	MaxOpenConn int
	MaxIdleConn int
}

type RedisConfig struct {
	Host     string
	Auth     string
	Database int
}

func Init() {
	NewViper := viper.NewWithOptions(viper.KeyDelimiter("::"))
	NewViper.SetConfigName("config")
	NewViper.AddConfigPath(".")
	NewViper.SetConfigType("yaml")
	NewViper.AutomaticEnv()
	NewViper.AllowEmptyEnv(true)

	if err := NewViper.ReadInConfig(); err != nil {
		log.Fatal().Err(err).Msg("config: viper reading config file failed")
	}

	err := NewViper.Unmarshal(&Config)
	if err != nil {
		log.Fatal().Err(err).Msg("config: viper decode config failed")
		return
	}
}
