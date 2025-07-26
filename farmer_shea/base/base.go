package base

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sheawinkler/farmer-shea/base/erc20"
	"github.com/sheawinkler/farmer-shea/base/nonfungiblepositionmanager"
	"github.com/sheawinkler/farmer-shea/base/uniswapv3factory"
)

const (
	UniswapV3FactoryAddress             = "0x33128a8fC17869897dcE68Ed026d694621f6FDfD"
	NonfungiblePositionManagerAddress = "0xC36442b4a4522E871399CD717aBDD847Ab11FE88"
)

// Client is a client for interacting with the Base blockchain.	
ype Client struct {
	client *ethclient.Client
}

// NewClient creates a new Base client.
func NewClient(rpcURL string) (*Client, error) {
	c, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, err
	}
	return &Client{client: c}, nil
}

// GetUniswapV3PoolAddress returns the address of a Uniswap V3 pool.
func (c *Client) GetUniswapV3PoolAddress(tokenA, tokenB common.Address, fee *big.Int) (common.Address, error) {
	factoryAddress := common.HexToAddress(UniswapV3FactoryAddress)
	factory, err := uniswapv3factory.NewUniswapv3factory(factoryAddress, c.client)
	if err != nil {
		return common.Address{}, err
	}

	poolAddress, err := factory.GetPool(nil, tokenA, tokenB, fee)
	if err != nil {
		return common.Address{}, err
	}

	return poolAddress, nil
}

// AddLiquidity adds liquidity to a Uniswap V3 pool.
func (c *Client) AddLiquidity(privateKey *ecdsa.PrivateKey, params nonfungiblepositionmanager.INonfungiblePositionManagerMintParams) error {
	npm, err := nonfungiblepositionmanager.NewNonfungiblepositionmanager(common.HexToAddress(NonfungiblePositionManagerAddress), c.client)
	if err != nil {
		return err
	}

	// Create a new transactor
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := c.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8453))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Add liquidity
	_, err = npm.Mint(auth, params)
	return err
}

// Approve approves a token for spending by another address.
func (c *Client) Approve(privateKey *ecdsa.PrivateKey, tokenAddress, spenderAddress common.Address, amount *big.Int) error {
	token, err := erc20.NewErc20(tokenAddress, c.client)
	if err != nil {
		return err
	}

	// Create a new transactor
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	nonce, err := c.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(8453))
	if err != nil {
		return err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// Approve the token
	_, err = token.Approve(auth, spenderAddress, amount)
	return err
}

// Swap simulates a swap on Uniswap V3.
func (c *Client) Swap(poolAddress common.Address, amount *big.Int) error {
	fmt.Printf("Simulating swap of %s tokens on pool %s\n", amount.String(), poolAddress.Hex())
	// In a real implementation, this would involve creating and sending a transaction
	// to the Uniswap V3 router contract.
	return nil
}
