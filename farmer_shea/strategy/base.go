package strategy

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// BaseStrategy defines the interface for a Base farming strategy.	"ype BaseStrategy interface {
	Execute(client *base.Client, w wallet.Wallet) error
}

// --- Simple Yield Farming Strategy ---

type simpleYieldFarmingStrategy struct{}

func NewSimpleYieldFarmingStrategy() BaseStrategy {
	return &simpleYieldFarmingStrategy{}
}

func (s *simpleYieldFarmingStrategy) Execute(client *base.Client, w wallet.Wallet) error {
	fmt.Println("Executing simple yield farming strategy on Base...")

	// Example: Get a USDC-WETH pool with a 0.05% fee
	usdc := common.HexToAddress("0x833589fCD6eDbE023dEEd136f9aAd50C355A4dF7")
	weth := common.HexToAddress("0x4200000000000000000000000000000000000006")
	fee := big.NewInt(500) // 0.05%

	poolAddress, err := client.GetUniswapV3PoolAddress(usdc, weth, fee)
	if err != nil {
		return fmt.Errorf("failed to get Uniswap V3 pool address: %w", err)
	}

	fmt.Printf("Uniswap V3 Pool Address: %s\n", poolAddress.Hex())

	return nil
}