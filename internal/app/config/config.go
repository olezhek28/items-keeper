package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TelegramToken     string
	PocketConsumerKey string
	AuthServerURL     string
	TelegramBotURL    string `mapstructure:"telegram_bot_url"`
	DBPath            string `mapstructure:"db_file"`
}

func Init() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	err = parseEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseEnv(cfg *Config) error {
	err := viper.BindEnv("telegram_token")
	if err != nil {
		return err
	}

	err = viper.BindEnv("pocket_consumer_key")
	if err != nil {
		return err
	}

	err = viper.BindEnv("auth_server_url")
	if err != nil {
		return err
	}

	cfg.TelegramToken = viper.GetString("telegram_token")
	cfg.PocketConsumerKey = viper.GetString("pocket_consumer_key")
	cfg.AuthServerURL = viper.GetString("auth_server_url")

	return nil
}
