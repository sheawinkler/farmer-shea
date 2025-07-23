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

type simpleVaultDepositStrategy struct{}

func NewSimpleVaultDepositStrategy() HyperliquidStrategy {
	return &simpleVaultDepositStrategy{}
}

func (s *simpleVaultDepositStrategy) Execute(client *hyperliquid.Client, w wallet.Wallet) error {
	fmt.Println("Executing simple vault deposit strategy on Hyperliquid...")
	return client.DepositToVault("100", sampleVaultAddress)
}
