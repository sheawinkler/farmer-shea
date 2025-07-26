package config

import (
	"github.com/spf13/viper"
)

// MovingAverageCrossoverConfig holds configuration for the MovingAverageCrossoverStrategy.
type MovingAverageCrossoverConfig struct {
	ShortPeriod int    `mapstructure:"short_period"`
	LongPeriod  int    `mapstructure:"long_period"`
	Symbol      string `mapstructure:"symbol"`
}

// Config is the configuration for the application.
type Config struct {
	WalletPath                 string                       `mapstructure:"wallet_path"`
	SolanaRPC                  string                       `mapstructure:"solana_rpc"`
	BaseRPC                    string                       `mapstructure:"base_rpc"`
	MovingAverageCrossoverConf MovingAverageCrossoverConfig `mapstructure:"moving_average_crossover"`
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