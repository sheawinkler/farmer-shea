package util

import (
	"math"

	"github.com/sonirico/go-hyperliquid"
)

// CalculateSMA calculates the Simple Moving Average.
func CalculateSMA(klines []hyperliquid.Kline, period int) float64 {
	if len(klines) < period {
		return 0.0
	}

	sum := 0.0
	for i := len(klines) - period; i < len(klines); i++ {
		sum += klines[i].Close
	}
	return sum / float64(period)
}

// CalculateStdDev calculates the standard deviation of a series of klines.
func CalculateStdDev(klines []hyperliquid.Kline, period int) float64 {
	if len(klines) < period {
		return 0.0
	}

	mean := CalculateSMA(klines, period)

	sum := 0.0
	for i := len(klines) - period; i < len(klines); i++ {
		sum += math.Pow(klines[i].Close-mean, 2)
	}

	return math.Sqrt(sum / float64(period))
}
