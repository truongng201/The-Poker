package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConfigType struct {
	// Database
	DatabaseURI string
	// Server environment
	Environment string
	// Allowed origins
	AllowedOrigins []string
	// JWT
	JWT JWTConfig
	// Redis
	Redis RedisConfig
	// Email
	Email EmailConfig
}

type JWTConfig struct {
	SecretKey                  string
	AccessTokenExpirationTime  int
	RefreshTokenExpirationTime int
}

type RedisConfig struct {
	Address      string
	Password     string
	DB           int
	PoolSize     int
	MinIdleConns int
	IdleTimeout  int
	MaxConnAge   int
	ReadTimeout  int
	WriteTimeout int
}

type EmailConfig struct {
	Username string
	Password string
	From     string
	FromName string
}

var Con ConfigType

func LoadConfig() (config ConfigType, err error) {
	log.Info("Loading config file")

	Config := &Con
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(filepath.Join(rootDir(), "config"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Error(fmt.Sprintf("Error reading config file, %s", err))
	}

	if err := viper.Unmarshal(&Config); err != nil {
		log.Error(fmt.Sprintf("Unable to decode into struct, %v", err))
	}
	return
}

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}
