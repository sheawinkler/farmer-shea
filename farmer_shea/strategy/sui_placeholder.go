package strategy

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/sheawinkler/farmer-shea/sui"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// --- Sui Placeholder Strategy ---

type suiPlaceholderStrategy struct {
	suiClient *sui.Client
}

func NewSuiPlaceholderStrategy(client *sui.Client) Strategy {
	return &suiPlaceholderStrategy{suiClient: client}
}

func (s *suiPlaceholderStrategy) Name() string {
	return "SuiPlaceholder"
}

func (s *suiPlaceholderStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	fmt.Println("Executing placeholder strategy on Sui...")
	// This is a placeholder. A real implementation would involve:
	// 1. Getting a pool address.
	// 2. Approving the router to spend tokens.
	// 3. Adding liquidity to the pool.
	return nil
}
