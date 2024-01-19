package config

import (
	"fmt"
	"github.com/jinglanghe/go-start/utils/log"
	"github.com/spf13/viper"
	"time"
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
	Database   DbStruct
	Redis      RedisConfig
}

type ApiPrefix struct {
	Group string
}

// RPCClient is RPC client config.
type RPCClient struct {
	Host    map[string]string
	Dial    time.Duration
	Timeout time.Duration
}

// RPCServer is RPC server config.
type RPCServer struct {
	Network           string
	Addr              string
	Timeout           time.Duration
	IdleTimeout       time.Duration
	MaxLifeTime       time.Duration
	ForceCloseWait    time.Duration
	KeepAliveInterval time.Duration
	KeepAliveTimeout  time.Duration
}

// HTTPServer is http server config.
type HTTPServer struct {
	Network      string
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HttpClient is http client config.
type HttpClient struct {
	MaxIdleConn        int
	MaxConnPerHost     int
	MaxIdleConnPerHost int
	TimeoutSeconds     int
}

type DbStruct struct {
	ServerType      string
	Username        string
	Password        string
	Host            string
	Port            int
	DbName          string
	SslMode         string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifetime int
	DbPing          int
	Debug           bool
}

func (d *DbStruct) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		d.Host, d.Port, d.Username, d.DbName, d.Password, d.SslMode)
}

type RedisConfig struct {
	Addr     string
	Password string
	Expire   time.Duration
	Db       int
}

func Init() *AppConfig {
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
	}

	return &Config
}
