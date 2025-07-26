package strategy

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// --- Uniswap V3 LP Strategy ---

type uniswapV3LPStrategy struct {
	baseClient *base.Client
	tokenA     common.Address
	tokenB     common.Address
	fee        *big.Int
	amountA    *big.Int
	amountB    *big.Int
}

func NewUniswapV3LPStrategy(client *base.Client, tokenA, tokenB, amountA, amountB string, fee int64) Strategy {
	a, _ := new(big.Int).SetString(amountA, 10)
	b, _ := new(big.Int).SetString(amountB, 10)
	return &uniswapV3LPStrategy{
		baseClient: client,
		tokenA:     common.HexToAddress(tokenA),
		tokenB:     common.HexToAddress(tokenB),
		fee:        big.NewInt(fee),
		amountA:    a,
		amountB:    b,
	}
}

func (s *uniswapV3LPStrategy) Name() string {
	return "UniswapV3LP"
}

func (s *uniswapV3LPStrategy) Execute(w wallet.Wallet) error {
	fmt.Println("Executing Uniswap V3 LP strategy on Base...")

	poolAddress, err := s.baseClient.GetUniswapV3PoolAddress(s.tokenA, s.tokenB, s.fee)
	if err != nil {
		return fmt.Errorf("failed to get Uniswap V3 pool address: %w", err)
	}

	fmt.Printf("Uniswap V3 Pool Address: %s\n", poolAddress.Hex())

	// This is a placeholder. A real implementation would involve:
	// 1. Approving the router to spend tokens.
	// 2. Adding liquidity to the pool.

	fmt.Printf("Simulating adding liquidity to pool %s with amounts %s and %s\n", poolAddress.Hex(), s.amountA.String(), s.amountB.String())

	return nil
}
