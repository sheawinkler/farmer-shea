package hyperliquid

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sonirico/go-hyperliquid"
)

// Client is a client for interacting with the Hyperliquid API.	ype Client struct {
	client *hyperliquid.Client
}

// NewClient creates a new Hyperliquid client.
func NewClient() (*Client, error) {
	c, err := hyperliquid.New(nil)
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
func (c *Client) GetKlines(symbol string, interval string, limit int) ([]hyperliquid.Kline, error) {
	url := fmt.Sprintf("https://api.hyperliquid.xyz/info/klines?coin=%s&interval=%s&limit=%d", symbol, interval, limit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var klines []hyperliquid.Kline
	if err := json.Unmarshal(body, &klines); err != nil {
		return nil, err
	}

	return klines, nil
}

// VaultDetails represents the details of a Hyperliquid vault.
type VaultDetails struct {
	Asset  string  `json:"asset"`
	Apy    float64 `json:"apy"`
	Tvl    float64 `json:"tvl"`
}

// GetVaultDetails fetches the details for a given vault address.
func (c *Client) GetVaultDetails(vaultAddress string) (*VaultDetails, error) {
	url := "https://api.hyperliquid.xyz/info"
	data := []byte(fmt.Sprintf(`{"type": "vaultDetails", "vaultAddress": "%s"}`, vaultAddress))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var vaultDetails VaultDetails
	if err := json.Unmarshal(body, &vaultDetails); err != nil {
		return nil, err
	}

	return &vaultDetails, nil
}
