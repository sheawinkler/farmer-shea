package main

import (
	"fmt"
	"os"

	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/execution"
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
		appUI.Log("Starting Farmer Shea Bot...\n")

		// Load or create a wallet
		w, err := wallet.Load(walletPath)
		if err != nil {
			if os.IsNotExist(err) {
				appUI.Log("No wallet found, creating a new one...\n")
				w, err = wallet.NewWallet()
				if err != nil {
					appUI.Log(fmt.Sprintf("Failed to create a new wallet: %v\n", err))
					return
				}
				if err := w.Save(walletPath); err != nil {
					appUI.Log(fmt.Sprintf("Failed to save the new wallet: %v\n", err))
					return
				}
				appUI.Log("New wallet created and saved.\n")
			} else {
				appUI.Log(fmt.Sprintf("Failed to load wallet: %v\n", err))
				return
			}
		}
		appUI.Log(fmt.Sprintf("Loaded wallet with public key: %s\n", w.PublicKey.String()))

		// Initialize Solana client
		solanaClient, err := solana.NewClient(rpc.MainNetBeta_RPC)
		if err != nil {
			appUI.Log(fmt.Sprintf("Failed to create Solana client: %v\n", err))
			return
		}

		// Get latest block height
		blockHeight, err := solanaClient.GetLatestBlockHeight()
		if err != nil {
			appUI.Log(fmt.Sprintf("Failed to get latest block height: %v\n", err))
			return
		}
		appUI.Log(fmt.Sprintf("Successfully connected to Solana. Latest block height: %d\n", blockHeight))

		// Initialize Hyperliquid client
		hyperliquidClient, err := hyperliquid.NewClient()
		if err != nil {
			appUI.Log(fmt.Sprintf("Failed to create Hyperliquid client: %v\n", err))
			return
		}

		// Initialize Base client
		baseClient, err := base.NewClient("https://mainnet.base.org")
		if err != nil {
			appUI.Log(fmt.Sprintf("Failed to create Base client: %v\n", err))
			return
		}

		// Execute strategies
		executor := execution.NewSimpleExecutor()

		solendStrategy := strategy.NewSolend(solanaClient, oracle.NewPyth(solanaClient.RPC()))
		if err := executor.Execute(solendStrategy, w); err != nil {
			appUI.Log(fmt.Sprintf("Failed to execute Solend strategy: %v\n", err))
		}

		hyperliquidStrategy := strategy.NewSimpleVaultDepositStrategy()
		if err := hyperliquidStrategy.Execute(hyperliquidClient, w); err != nil {
			appUI.Log(fmt.Sprintf("Failed to execute Hyperliquid strategy: %v\n", err))
		}

		baseStrategy := strategy.NewSimpleYieldFarmingStrategy()
		if err := baseStrategy.Execute(baseClient, w); err != nil {
			appUI.Log(fmt.Sprintf("Failed to execute Base strategy: %v\n", err))
		}

		appUI.Log("Farmer Shea Bot cycle complete.\n")
	}()

	if err := appUI.Run(); err != nil {
		panic(err)
	}
}