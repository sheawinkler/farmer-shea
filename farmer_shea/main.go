package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/config"
	"github.com/sheawinkler/farmer-shea/executor"
	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/oracle"
	"github.com/sheawinkler/farmer-shea/solana"
	"github.com/sheawinkler/farmer-shea/strategy"
	"github.com/sheawinkler/farmer-shea/ui"
	"github.com/sheawinkler/farmer-shea/wallet"

	"github.com/gagliardetto/solana-go/rpc"
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
		pnl := 123.45

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
		solanaClient, err := solana.NewClient(cfg.SolanaRPC)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create Solana client")
		}

		// Get latest block height
		blockHeight, err := solanaClient.GetLatestBlockHeight()
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to get latest block height")
		}
		log.Info().Uint64("blockHeight", blockHeight).Msg("Successfully connected to Solana")

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

		// Execute strategies
		executor := executor.NewSimpleExecutor()

		solendStrategy := strategy.NewSolend(solanaClient, oracle.NewPyth(solanaClient.RPC()))
		if err := executor.Execute(solendStrategy, w); err != nil {
			log.Error().Err(err).Msg("Failed to execute Solend strategy")
		}

		hyperliquidStrategy := strategy.NewSimpleVaultDepositStrategy()
		if err := hyperliquidStrategy.Execute(hyperliquidClient, w); err != nil {
			log.Error().Err(err).Msg("Failed to execute Hyperliquid strategy")
		}

		baseStrategy := strategy.NewSimpleYieldFarmingStrategy()
		if err := baseStrategy.Execute(baseClient, w); err != nil {
			log.Error().Err(err).Msg("Failed to execute Base strategy")
		}

		log.Info().Msg("Farmer Shea Bot cycle complete.")
	}()

	if err := appUI.Run(); err != nil {
		log.Fatal().Err(err).Msg("UI error")
	}
}
