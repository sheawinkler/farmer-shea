package executor

import (
	"crypto/ecdsa"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/sheawinkler/farmer-shea/strategy"
	"github.com/sheawinkler/farmer-shea/wallet"
)

const (
	defaultRetryAttempts = 3
	defaultRetryDelay    = 5 * time.Second
)

// Executor manages the execution of strategies.
type Executor struct {
	strategies []strategy.Strategy
	privateKey *ecdsa.PrivateKey
	wallet     wallet.Wallet
}

// New creates a new Executor.
func New(strategies []strategy.Strategy, w wallet.Wallet, pk *ecdsa.PrivateKey) *Executor {
	return &Executor{
		strategies: strategies,
		wallet:     w,
		privateKey: pk,
	}
}

// Run starts the execution of all strategies.
func (e *Executor) Run() {
	for _, s := range e.strategies {
		go e.runStrategy(s)
	}
}

func (e *Executor) runStrategy(s strategy.Strategy) {
	for {
		log.Info().Str("strategy", s.Name()).Msg("Executing strategy")
		var err error
		for i := 0; i < defaultRetryAttempts; i++ {
			err = s.Execute(e.wallet, e.privateKey)
			if err == nil {
				break
			}
			log.Error().Err(err).Str("strategy", s.Name()).Int("attempt", i+1).Msg("Error executing strategy, retrying...")
			time.Sleep(defaultRetryDelay)
		}

		if err != nil {
			log.Error().Err(err).Str("strategy", s.Name()).Msg("Strategy execution failed after multiple attempts")
		}

		time.Sleep(5 * time.Minute) // Execute every 5 minutes
	}
}
