package sui

import (
	"context"

	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
)

// Client is a Sui client.
type Client struct {
	*client.Client
}

// NewClient creates a new Sui client.
func NewClient(rpcEndpoint string) (*Client, error) {
	c, err := client.Dial(rpcEndpoint)
	if err != nil {
		return nil, err
	}
	return &Client{c}, nil
}

// GetLatestBlockHeight gets the latest block height of the Sui blockchain.
func (c *Client) GetLatestBlockHeight() (uint64, error) {
	return c.GetLatestCheckpointSequenceNumber(context.Background())
}
