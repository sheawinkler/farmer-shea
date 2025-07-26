package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/rs/zerolog/log"
	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/config"
	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/solana"
	"github.com/sheawinkler/farmer-shea/strategy"
	"github.com/sheawinkler/farmer-shea/ui"
	"github.com/sheawinkler/farmer-shea/wallet"
)

const walletPath = "farmer_shea_wallet.json"

func main() {
	appUI := ui.New()

	go func() {
		log.Info().Msg("Starting Farmer Shea Bot...")

		// Load config
		cfg, err := config.Load()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to load config")
		}

		// Dummy data for portfolio and PnL
		portfolio := map[string]float64{
			"SOL": 10.5,
			"ETH": 2.3,
			"BTC": 0.5,
		}
		pnl := map[string]float64{
			"SOL": 123.45,
			"ETH": 67.89,
			"BTC": -12.34,
		}

		appUI.UpdatePortfolio(portfolio)
		appUI.UpdatePnL(pnl)

		// Load or create a wallet
		w, err := wallet.Load(cfg.WalletPath)
		if err != nil {
			if os.IsNotExist(err) {
				log.Info().Msg("No wallet found, creating a new one...")
				w, err = wallet.NewWallet()
				if err != nil {
					log.Fatal().Err(err).Msg("Failed to create a new wallet")
				}
				if err := w.Save(cfg.WalletPath); err != nil {
					log.Fatal().Err(err).Msg("Failed to save the new wallet")
				}
				log.Info().Msg("New wallet created and saved.")
			} else {
				log.Fatal().Err(err).Msg("Failed to load wallet")
			}
		}
		log.Info().Str("publicKey", w.PublicKey.String()).Msg("Loaded wallet")

		// Initialize Solana client
		_, err = solana.NewClient(cfg.SolanaRPC)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create Solana client")
		}

		// Initialize Hyperliquid client
		hyperliquidClient, err := hyperliquid.NewClient()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create Hyperliquid client")
		}

		// Initialize Base client
		baseClient, err := base.NewClient(cfg.BaseRPC)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create Base client")
		}

		// Initialize Strategy Manager
		strategyManager := strategy.NewManager()

		// Add Strategies
		strategyManager.Add(strategy.NewSimpleVaultDepositStrategy(hyperliquidClient, cfg.Hyperliquid.VaultAddress, cfg.Hyperliquid.Amount))
		strategyManager.Add(strategy.NewUniswapV3LPStrategy(baseClient, cfg.Base.TokenA, cfg.Base.TokenB, cfg.Base.AmountA, cfg.Base.AmountB, cfg.Base.Fee))

		// Run the strategies
		for _, s := range strategyManager.Strategies {
			go func(s strategy.Strategy) {
				for {
					log.Info().Str("strategy", s.Name()).Msg("Executing strategy")
					if err := s.Execute(*w, w.PrivateKey); err != nil {
						log.Error().Err(err).Str("strategy", s.Name()).Msg("Error executing strategy")
					}
					time.Sleep(5 * time.Minute) // Execute every 5 minutes
				}
			}(s)
		}

		log.Info().Msg("Farmer Shea Bot cycle complete.")
	}()

	if err := appUI.Run(); err != nil {
		log.Fatal().Err(err).Msg("UI error")
	}
}