package config

import (
	"context"

	"backend-service/internal/infrastructure/database"

	"github.com/sethvargo/go-envconfig"
)

var conf *AppConfig

type AppConfig struct {
	Client *database.Config
	// GrowthBookUrl     string `env:"THIRDPARTY_GROWTHBOOK_URL"`
	// AuthBaseURL       string `env:"AUTH_BASE_URL"`
	Port     string `env:"API_PORT"`
	Env      string `env:"APP_ENV"`
	Basepath string `env:"BASE_PATH"`
	// OmsSyncItemApi    string `env:"OMS_SYNC_ITEM_API"`
	// OmsSyncItemApiKey string `env:"OMS_SYNC_ITEM_API_KEY"`
}

func GetAppconfig() *AppConfig {
	var config AppConfig
	if err := envconfig.Process(context.Background(), &config); err != nil {
		panic(err)
	}

	conf = &config

	return conf
}
