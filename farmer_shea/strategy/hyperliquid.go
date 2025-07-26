package strategy

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// --- Simple Vault Deposit Strategy ---

type simpleVaultDepositStrategy struct {
	hyperliquidClient *hyperliquid.Client
	vaultAddress      string
	amount            string
}

func NewSimpleVaultDepositStrategy(client *hyperliquid.Client, vaultAddress string, amount string) Strategy {
	return &simpleVaultDepositStrategy{
		hyperliquidClient: client,
		vaultAddress:      vaultAddress,
		amount:            amount,
	}
}

func (s *simpleVaultDepositStrategy) Name() string {
	return "SimpleVaultDeposit"
}

func (s *simpleVaultDepositStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	fmt.Printf("Executing simple vault deposit strategy on Hyperliquid with amount %s to vault %s...\n", s.amount, s.vaultAddress)
	return s.hyperliquidClient.DepositToVault(s.amount, s.vaultAddress)
}
