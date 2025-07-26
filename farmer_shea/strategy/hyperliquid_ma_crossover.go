package strategy

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/util"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// --- Moving Average Crossover Strategy ---

type maCrossoverStrategy struct {
	hyperliquidClient *hyperliquid.Client
	symbol            string
	shortPeriod       int
	longPeriod        int
}

func NewMACrossoverStrategy(client *hyperliquid.Client, symbol string, shortPeriod, longPeriod int) Strategy {
	return &maCrossoverStrategy{
		hyperliquidClient: client,
		symbol:            symbol,
		shortPeriod:       shortPeriod,
		longPeriod:        longPeriod,
	}
}

func (s *maCrossoverStrategy) Name() string {
	return "MACrossover"
}

func (s *maCrossoverStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	klines, err := s.hyperliquidClient.GetKlines(s.symbol, "1h", s.longPeriod)
	if err != nil {
		return err
	}

	shortSMA := util.CalculateSMA(klines, s.shortPeriod)
	longSMA := util.CalculateSMA(klines, s.longPeriod)

	fmt.Printf("Executing MA Crossover strategy for %s. Short SMA: %f, Long SMA: %f\n", s.symbol, shortSMA, longSMA)

	if shortSMA > longSMA {
		fmt.Println("BUY SIGNAL")
		// In a real implementation, this would involve placing a buy order.
	} else if shortSMA < longSMA {
		fmt.Println("SELL SIGNAL")
		// In a real implementation, this would involve placing a sell order.
	}

	return nil
}
