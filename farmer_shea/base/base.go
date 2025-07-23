package base

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client is a client for interacting with the Base blockchain.	ype Client struct {
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
