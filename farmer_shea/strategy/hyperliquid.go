package strategy

import (
	"fmt"

	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/wallet"
)

const (
	// This is a sample vault address. In a real-world scenario, we would want to have a more sophisticated way of selecting a vault.
	sampleVaultAddress = "0x1234567890123456789012345678901234567890"
)

// HyperliquidStrategy defines the interface for a Hyperliquid farming strategy.	ype HyperliquidStrategy interface {
	Execute(client *hyperliquid.Client, w wallet.Wallet) error
}

// --- Simple Vault Deposit Strategy ---

type simpleVaultDepositStrategy struct {
	hyperliquidClient *hyperliquid.Client
}

func NewSimpleVaultDepositStrategy(client *hyperliquid.Client) HyperliquidStrategy {
	return &simpleVaultDepositStrategy{hyperliquidClient: client}
}

func (s *simpleVaultDepositStrategy) Execute(w wallet.Wallet) error {
	fmt.Println("Executing simple vault deposit strategy on Hyperliquid...")
	return s.hyperliquidClient.DepositToVault("100", sampleVaultAddress)
}
