package hyperliquid

import (
	"fmt"

	"github.com/sonirico/go-hyperliquid/pkg/client"
	"github.com/sonirico/go-hyperliquid/pkg/exchange"
	"github.com/sonirico/go-hyperliquid/pkg/ws"
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

// GetKlines fetches historical klines for a given symbol and interval.
func (c *Client) GetKlines(symbol string, interval string, limit int) ([]ws.Kline, error) {
	// The go-hyperliquid library's client.Exchange does not directly expose a klines endpoint.
	// This would typically require direct HTTP requests to the Hyperliquid API or extending the SDK.
	// For now, we'll return a dummy response.
	return []ws.Kline{},
		fmt.Errorf("GetKlines not yet implemented for Hyperliquid client")
}
