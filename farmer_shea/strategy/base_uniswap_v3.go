package strategy

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/base/nonfungiblepositionmanager"
	"github.com/sheawinkler/farmer-shea/util"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// --- Uniswap V3 LP Strategy ---

type uniswapV3LPStrategy struct {
	baseClient *base.Client
	hyperliquidClient *hyperliquid.Client
	tokenA     common.Address
	tokenB     common.Address
	fee        *big.Int
	amountA    *big.Int
	amountB    *big.Int
}

func NewUniswapV3LPStrategy(client *base.Client, hyperliquidClient *hyperliquid.Client, tokenA, tokenB, amountA, amountB string, fee int64) Strategy {
	a, _ := new(big.Int).SetString(amountA, 10)
	b, _ := new(big.Int).SetString(amountB, 10)
	return &uniswapV3LPStrategy{
		baseClient: client,
		hyperliquidClient: hyperliquidClient,
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

func (s *uniswapV3LPStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	fmt.Println("Executing Uniswap V3 LP strategy on Base...")

	// Approve the router to spend tokens
	if err := s.baseClient.Approve(privateKey, s.tokenA, common.HexToAddress(base.NonfungiblePositionManagerAddress), s.amountA); err != nil {
		return fmt.Errorf("failed to approve token A: %w", err)
	}
	if err := s.baseClient.Approve(privateKey, s.tokenB, common.HexToAddress(base.NonfungiblePositionManagerAddress), s.amountB); err != nil {
		return fmt.Errorf("failed to approve token B: %w", err)
	}

	// Calculate the tick range
	tickLower, tickUpper, err := s.calculateTickRange()
	if err != nil {
		return err
	}

	// Add liquidity to the pool
	params := nonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
		Token0:         s.tokenA,
		Token1:         s.tokenB,
		Fee:            s.fee,
		TickLower:      tickLower,
		TickUpper:      tickUpper,
		Amount0Desired: s.amountA,
		Amount1Desired: s.amountB,
		Amount0Min:     big.NewInt(0),
		Amount1Min:     big.NewInt(0),
		Recipient:      w.PublicKey,
		Deadline:       big.NewInt(time.Now().Add(15 * time.Minute).Unix()),
	}

	return s.baseClient.AddLiquidity(privateKey, params)
}

func (s *uniswapV3LPStrategy) calculateTickRange() (*big.Int, *big.Int, error) {
	// This is a simplified implementation. A more robust implementation would
	// involve using a more sophisticated volatility model.
	klines, err := s.hyperliquidClient.GetKlines("ETH", "1h", 100)
	if err != nil {
		return nil, nil, err
	}

	stdDev := util.CalculateStdDev(klines, 100)

	// This is a simplified way to calculate the tick range. A more robust
	// implementation would involve using a more sophisticated formula.
	tickLower := new(big.Int).SetInt64(int64(klines[len(klines)-1].Close - stdDev))
	tickUpper := new(big.Int).SetInt64(int64(klines[len(klines)-1].Close + stdDev))

	return tickLower, tickUpper, nil
}
