package solana

import (
	"context"

	"github.com/gagliardetto/solana-go/rpc"
)

// Client is a Solana client.
type Client struct {
	*rpc.Client
}

// NewClient creates a new Solana client.
func NewClient(rpcEndpoint string) (*Client, error) {
	client := rpc.New(rpcEndpoint)
	return &Client{client}, nil
}

// GetLatestBlockHeight gets the latest block height of the Solana blockchain.
func (c *Client) GetLatestBlockHeight() (uint64, error) {
	return c.GetSlot(context.Background(), rpc.CommitmentFinalized)
}
