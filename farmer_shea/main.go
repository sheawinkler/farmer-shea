package main

import (
	"farmer_shea/execution"
	"farmer_shea/oracle"
	"farmer_shea/solana"
	"farmer_shea/strategy"
	"farmer_shea/wallet"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

const walletPath = "farmer_shea_wallet.json"

func main() {
	// Set up logging
	logFile, err := os.OpenFile("farmer_shea.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	log.Println("Starting Farmer Shea Bot...")

	// Load or create a wallet
	w, err := wallet.Load(walletPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("No wallet found, creating a new one...")
			w, err = wallet.NewWallet()
			if err != nil {
				log.Fatalf("Failed to create a new wallet: %v", err)
			}
			if err := w.Save(walletPath); err != nil {
				log.Fatalf("Failed to save the new wallet: %v", err)
			}
			log.Println("New wallet created and saved.")
		} else {
			log.Fatalf("Failed to load wallet: %v", err)
		}
	}
	log.Printf("Loaded wallet with public key: %s", w.PublicKey.String())

	// Initialize Solana client
	solanaClient, err := solana.NewClient(rpc.MainNetBeta_RPC)
	if err != nil {
		log.Fatalf("Failed to create Solana client: %v", err)
	}

	// Get latest block height
	blockHeight, err := solanaClient.GetLatestBlockHeight()
	if err != nil {
		log.Fatalf("Failed to get latest block height: %v", err)
	}
	log.Printf("Successfully connected to Solana. Latest block height: %d", blockHeight)

	// Get token prices from Jupiter
	mints := []string{"So11111111111111111111111111111111111111112", "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v"}
	prices, err := oracle.GetJupiterPrices(mints)
	if err != nil {
		// Don't fail the whole process if Jupiter is down
		log.Printf("Failed to get prices from Jupiter: %v", err)
	} else {
		log.Printf("Prices from Jupiter: %v", prices)
	}

	// Get markets from Solend
	markets, err := oracle.GetSolendMarkets()
	if err != nil {
		log.Fatalf("Failed to get markets from Solend: %v", err)
	}
	log.Printf("Markets from Solend: %v", markets)

	// Determine the best lending market
	bestMarket, err := strategy.DetermineBestLendingMarket(markets)
	if err != nil {
		log.Fatalf("Failed to determine best lending market: %v", err)
	}
	log.Printf("Best lending market: %v", bestMarket)

	// Get user portfolio from Solend
	portfolio, err := oracle.GetSolendPortfolio(w.PublicKey.String())
	if err != nil {
		log.Fatalf("Failed to get portfolio from Solend: %v", err)
	}
	log.Printf("Portfolio from Solend: %v", portfolio)

	// Send a test transaction
	recipientWallet, err := wallet.NewWallet()
	if err != nil {
		log.Fatalf("Failed to create a new wallet for recipient: %v", err)
	}
	recipient := recipientWallet.PublicKey
	log.Printf("Attempting to send 1 lamport to %s. This is expected to fail as the wallet has no funds.", recipient)
	sig, err := execution.SendSol(solanaClient.Client, w.PrivateKey, recipient, 1)
	if err != nil {
		log.Printf("Transaction failed as expected: %v", err)
	} else {
		log.Printf("Transaction succeeded unexpectedly with signature: %s", sig)
	}

	log.Println("Farmer Shea Bot cycle complete.")
}

