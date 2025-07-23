package hyperliquid

import (
	"github.com/sonirico/go-hyperliquid/pkg/client"
	"github.com/sonirico/go-hyperliquid/pkg/exchange"
)

// Client is a client for interacting with the Hyperliquid API.	ype Client struct {
	client *client.Client
}

// NewClient creates a new Hyperliquid client.
func NewClient() (*Client, error) {
	c, err := client.New(nil)
	if err != nil {
		return nil, err
	}
	return &Client{client: c}, nil
}

// DepositToVault deposits the given amount to the given vault.
func (c *Client) DepositToVault(amount string, vaultAddress string) error {
	_, err := c.client.Exchange.VaultDeposit(amount, vaultAddress)
	return err
}