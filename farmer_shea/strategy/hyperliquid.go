package strategy

import (
	"crypto/ecdsa"
	"fmt"
	"sort"

	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// --- Simple Vault Deposit Strategy ---

type simpleVaultDepositStrategy struct {
	hyperliquidClient *hyperliquid.Client
	amount            string
	stopLoss          float64
}

func NewSimpleVaultDepositStrategy(client *hyperliquid.Client, amount string, stopLoss float64) Strategy {
	return &simpleVaultDepositStrategy{
		hyperliquidClient: client,
		amount:            amount,
		stopLoss:          stopLoss,
	}
}

func (s *simpleVaultDepositStrategy) Name() string {
	return "SimpleVaultDeposit"
}

func (s *simpleVaultDepositStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	vaults, err := s.getVaults()
	if err != nil {
		return err
	}

	bestVault, err := s.determineBestVault(vaults)
	if err != nil {
		return err
	}

	if bestVault.Apy < s.stopLoss {
		fmt.Printf("Vault APY (%f) is below stop-loss threshold (%f). Withdrawing funds...\n", bestVault.Apy, s.stopLoss)
		return s.hyperliquidClient.WithdrawFromVault(s.amount, bestVault.Address)
	}

	fmt.Printf("Executing simple vault deposit strategy on Hyperliquid with amount %s to vault %s...\n", s.amount, bestVault.Address)
	return s.hyperliquidClient.DepositToVault(s.amount, bestVault.Address)
}

func (s *simpleVaultDepositStrategy) getVaults() ([]hyperliquid.VaultDetails, error) {
	// In a real implementation, this list would be fetched from a community-maintained list or by scraping the Hyperliquid UI.
	vaultAddresses := []string{
		"0x1a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0t",
		"0x1a2b3c4d5e6f7g8h9i0j1k2l3m4n5o6p7q8r9s0u",
	}

	var vaults []hyperliquid.VaultDetails
	for _, address := range vaultAddresses {
		details, err := s.hyperliquidClient.GetVaultDetails(address)
		if err != nil {
			// Log the error, but continue to the next vault
			fmt.Printf("Error getting details for vault %s: %s\n", address, err)
			continue
		}
		vaults = append(vaults, *details)
	}

	return vaults, nil
}

func (s *simpleVaultDepositStrategy) determineBestVault(vaults []hyperliquid.VaultDetails) (*hyperliquid.VaultDetails, error) {
	if len(vaults) == 0 {
		return nil, fmt.Errorf("no vaults found")
	}

	// Sort by APY in descending order
	sort.Slice(vaults, func(i, j int) bool {
		return vaults[i].Apy > vaults[j].Apy
	})

	// In a real implementation, we would also consider TVL and other factors.
	return &vaults[0], nil
}
