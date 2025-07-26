package config

import (
	"github.com/spf13/viper"
)

// HyperliquidConfig holds configuration for the Hyperliquid vault strategy.
type HyperliquidConfig struct {
	VaultAddress string  `mapstructure:"vault_address"`
	Amount       string  `mapstructure:"amount"`
	StopLoss     float64 `mapstructure:"stop_loss"`
}

// BaseConfig holds configuration for the Base Uniswap V3 LP strategy.
type BaseConfig struct {
	TokenA  string `mapstructure:"token_a"`
	TokenB  string `mapstructure:"token_b"`
	Fee     int64  `mapstructure:"fee"`
	AmountA string `mapstructure:"amount_a"`
	AmountB string `mapstructure:"amount_b"`
}

// Config is the configuration for the application.
type Config struct {
	WalletPath        string            `mapstructure:"wallet_path"`
	SolanaRPC         string            `mapstructure:"solana_rpc"`
	BaseRPC           string            `mapstructure:"base_rpc"`
	Hyperliquid       HyperliquidConfig `mapstructure:"hyperliquid"`
	Base              BaseConfig        `mapstructure:"base"`
}

// Load loads the configuration from a file.
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
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
