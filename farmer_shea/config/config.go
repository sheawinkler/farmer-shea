package config

import (
	"github.com/spf13/viper"
)

// Config is the configuration for the application.	ype Config struct {
	WalletPath      string `mapstructure:"wallet_path"`
	SolanaRPC       string `mapstructure:"solana_rpc"`
	BaseRPC         string `mapstructure:"base_rpc"`
}

// Load loads the configuration from a file.
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}