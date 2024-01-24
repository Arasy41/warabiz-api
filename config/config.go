package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server        ServerAccount
	Routes        RoutesAccount
	Logger        LoggerAccount
	Authorization AuthorizationAccount
	Cookies       CookiesAccount
	Connection    ConnectionAccount
}

type ServerAccount struct {
	Name          string
	AppVersion    string
	Port          string
	Env           string
	SSL           bool
	BodyLimit     int
	CSRF          bool
	HexaSecretKey string
}

type RoutesAccount struct {
	Methods        string
	Headers        string
	Origins        string
	DisableOrigins bool
}

type LoggerAccount struct {
	Encoding      string
	Level         string
	LogPath       string
	DisableCaller bool
	CallerSkipper int
	Logrotation   LogrotationAccount
	Logstash      LogstashAccount
	Loki          LokiAccount
}

type LogrotationAccount struct {
	RotationTime time.Duration
	MaxRotation  uint
}

type LogstashAccount struct {
	IsActive bool
	URI      string
}

type LokiAccount struct {
	IsActive bool
	URI      string
}

type AuthorizationAccount struct {
	Public    PublicAccount
	JWT       JWTAccount
	Signature SignatureAccount
	Session   SessionAccount
}

type PublicAccount struct {
	SecretKey string
}

type JWTAccount struct {
	AccessTokenSecretKey  string
	AccessTokenDuration   time.Duration
	RefreshTokenSecretKey string
	RefreshTokenDuration  time.Duration
	ShowTokenBody         bool
}

type SignatureAccount struct {
	AccessTokenDuration time.Duration
}

type SessionAccount struct {
	Name   string
	Prefix string
	Expire time.Duration
}

type CookiesAccount struct {
	CoreDomain string
	CoreAT     string
	CoreRT     string
}

type ConnectionAccount struct {
	Warabiz  DatabaseAccount
	Redis    RedisAccount
	MongoDB  MongoDBAccount
	SMTP     SMTPAccount
	RabbitMQ RabbitMQAccount
}

type DatabaseAccount struct {
	ServerType      string
	DriverSource    string
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	MaxIdleConns    int
	ConnMaxIdleTime time.Duration
}

type RedisAccount struct {
	Host         string
	Password     string
	DB           int
	DefaultDB    int
	MinIdleConns int
	PoolSize     int
	PoolTimeout  time.Duration
}

type MongoDBAccount struct {
	MongoURI string
}

type SMTPAccount struct {
	Host     string
	Port     int
	Username string
	Password string
	Timeout  time.Duration
	Sender   string
}

type RabbitMQAccount struct {
	Protocol string
	Host     string
	Port     int
	Username string
	Password string
}

//=================================================================================================================

// * Init Config
func InitConfig(env string) *Config {

	var config *viper.Viper
	var err error

	// if env == "prod" {
	// 	newConfig := viper.New()
	// 	buff := GetProductionConfig()
	// 	fmt.Println(string(buff))
	// 	err = newConfig.ReadConfig(bytes.NewBuffer(buff))
	// 	if err != nil {
	// 		log.Fatalf("ReadConfig: %v", err)
	// 	}
	// 	config = newConfig
	// } else {
	configPath := GetConfigPath(env)
	configFile, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	config = configFile
	// }

	cfg, err := ParseConfig(config)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	return cfg
}

// * Get config path for local or docker
func GetConfigPath(configPath string) string {
	if configPath == "staging" {
		return "./config/config-stg"
	}
	if configPath == "prod" {
		return "./config/config-prod"
	}
	return "./config/config-dev"
}

// * Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// * Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
