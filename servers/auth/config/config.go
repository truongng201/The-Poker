package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type config struct {
	// Database
	DatabaseURI string
	// Server environment
	Environment string
	// Allowed origins
	AllowedOrigins []string
	// JWT
	JWTSecretKey string
	// JWT Refresh Token Expiry Time
	JWTRefreshTokenExpireTime int
}

var Con config

func LoadConfig() (config config, err error) {
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
