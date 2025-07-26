package strategy

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sheawinkler/farmer-shea/base"
	"github.com/sheawinkler/farmer-shea/base/nonfungiblepositionmanager"
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

func (s *uniswapV3LPStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	fmt.Println("Executing Uniswap V3 LP strategy on Base...")

	// Approve the router to spend tokens
	if err := s.baseClient.Approve(privateKey, s.tokenA, common.HexToAddress(base.NonfungiblePositionManagerAddress), s.amountA); err != nil {
		return fmt.Errorf("failed to approve token A: %w", err)
	}
	if err := s.baseClient.Approve(privateKey, s.tokenB, common.HexToAddress(base.NonfungiblePositionManagerAddress), s.amountB); err != nil {
		return fmt.Errorf("failed to approve token B: %w", err)
	}

	// Add liquidity to the pool
	params := nonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
		Token0:         s.tokenA,
		Token1:         s.tokenB,
		Fee:            s.fee,
		TickLower:      big.NewInt(-887272),
		TickUpper:      big.NewInt(887272),
		Amount0Desired: s.amountA,
		Amount1Desired: s.amountB,
		Amount0Min:     big.NewInt(0),
		Amount1Min:     big.NewInt(0),
		Recipient:      w.PublicKey,
		Deadline:       big.NewInt(time.Now().Add(15 * time.Minute).Unix()),
	}

	return s.baseClient.AddLiquidity(params)
}
