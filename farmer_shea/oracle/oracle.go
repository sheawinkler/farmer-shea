package oracle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// PriceData represents the price information from CoinGecko.
type PriceData struct {
	Solana struct {
		USD float64 `json:"usd"`
	} `json:"solana"`
}

// JupiterPriceData represents the price information from Jupiter.
type JupiterPriceData struct {
	Data map[string]struct {
		Price float64 `json:"price"`
	} `json:"data"`
}

// SolendReserve represents a reserve on Solend.
type SolendReserve struct {
	Name      string `json:"name"`
	Liquidity struct {
		SupplyApy float64 `json:"supplyApy"`
	} `json:"liquidity"`
}

// SolendMarket represents a market on Solend.
type SolendMarket struct {
	Name      string `json:"name"`
	SupplyApy float64 `json:"supplyApy"`
}

// SolendMarkets represents a list of markets on Solend.
type SolendMarkets struct {
	Results []SolendMarket `json:"results"`
}

// SolendPortfolio represents a user's portfolio on Solend.
type SolendPortfolio struct {
	Deposits []struct {
		Amount float64 `json:"amount"`
		Symbol string  `json:"symbol"`
	} `json:"deposits"`
	Borrows []struct {
		Amount float64 `json:"amount"`
		Symbol string  `json:"symbol"`
	} `json:"borrows"`
}

// GetSolanaPrice gets the price of SOL in USD from CoinGecko.
func GetSolanaPrice() (float64, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=solana&vs_currencies=usd")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var priceData PriceData
	err = json.NewDecoder(resp.Body).Decode(&priceData)
	if err != nil {
		return 0, err
	}

	if priceData.Solana.USD == 0 {
		return 0, fmt.Errorf("could not fetch SOL price")
	}

	return priceData.Solana.USD, nil
}

// GetJupiterPrices gets the prices of a list of tokens in USD from Jupiter.
func GetJupiterPrices(mints []string) (map[string]float64, error) {
	url := fmt.Sprintf("https://price.jup.ag/v4/price?ids=%s", strings.Join(mints, ","))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var jupiterPriceData JupiterPriceData
	err = json.NewDecoder(resp.Body).Decode(&jupiterPriceData)
	if err != nil {
		return nil, err
	}

	if len(jupiterPriceData.Data) == 0 {
		return nil, fmt.Errorf("could not fetch prices from Jupiter")
	}

	prices := make(map[string]float64)
	for mint, data := range jupiterPriceData.Data {
		prices[mint] = data.Price
	}

	return prices, nil
}

// GetSolendMarkets gets a list of all markets on Solend.
func GetSolendMarkets() ([]SolendMarket, error) {
	resp, err := http.Get("https://api.solend.fi/v1/markets/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var solendMarkets SolendMarkets
	err = json.NewDecoder(resp.Body).Decode(&solendMarkets)
	if err != nil {
		return nil, err
	}

	return solendMarkets.Results, nil
}

// GetSolendPortfolio gets a user's portfolio from Solend.
func GetSolendPortfolio(walletAddress string) (*SolendPortfolio, error) {
	url := fmt.Sprintf("https://api.solend.fi/v1/users/%s/portfolio", walletAddress)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var portfolio SolendPortfolio
	err = json.NewDecoder(resp.Body).Decode(&portfolio)
	if err != nil {
		return nil, err
	}

	return &portfolio, nil
}
