package strategy

import (
	"fmt"
	"time"

	"github.com/sheawinkler/farmer-shea/hyperliquid"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// MovingAverageCrossoverStrategy is a strategy for perpetual futures trading on Hyperliquid.	
ype MovingAverageCrossoverStrategy struct {
	hyperliquidClient *hyperliquid.Client
	shortPeriod       int
	longPeriod        int
	symbol            string
}

// NewMovingAverageCrossoverStrategy creates a new MovingAverageCrossoverStrategy.
func NewMovingAverageCrossoverStrategy(client *hyperliquid.Client, symbol string, shortPeriod, longPeriod int) *MovingAverageCrossoverStrategy {
	return &MovingAverageCrossoverStrategy{
		hyperliquidClient: client,
		symbol:            symbol,
		shortPeriod:       shortPeriod,
		longPeriod:        longPeriod,
	}
}

func (s *MovingAverageCrossoverStrategy) Execute(w wallet.Wallet) error {
	fmt.Printf("Executing Moving Average Crossover strategy for %s...\n", s.symbol)

	// Fetch klines
	klines, err := s.hyperliquidClient.GetKlines(s.symbol, "1h", s.longPeriod+10) // Fetch enough data for both MAs
	if err != nil {
		return fmt.Errorf("failed to get klines: %w", err)
	}

	if len(klines) < s.longPeriod {
		return fmt.Errorf("not enough kline data for moving average calculation")
	}

	// Calculate moving averages
	shortMA := calculateSMA(klines, s.shortPeriod)
	longMA := calculateSMA(klines, s.longPeriod)

	fmt.Printf("Short MA (%d): %.2f, Long MA (%d): %.2f\n", s.shortPeriod, shortMA, s.longPeriod, longMA)

	// Determine signal (simplified for now)
	if shortMA > longMA {
		fmt.Println("Signal: Buy (Short MA > Long MA)")
		// TODO: Place long order
	} else if shortMA < longMA {
		fmt.Println("Signal: Sell (Short MA < Long MA)")
		// TODO: Place short order
	} else {
		fmt.Println("Signal: No clear trend")
	}

	return nil
}

// calculateSMA calculates the Simple Moving Average.
func calculateSMA(klines []ws.Kline, period int) float64 {
	if len(klines) < period {
		return 0.0
	}

	sum := 0.0
	for i := len(klines) - period; i < len(klines); i++ {
		sum += klines[i].Close
	}
	return sum / float64(period)
}
