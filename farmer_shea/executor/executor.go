package executor

import (
	"time"

	"github.com/sheawinkler/farmer-shea/strategy"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// Executor defines the interface for an executor.	ype Executor interface {
	Execute(strategy strategy.Strategy, wallet wallet.Wallet) error
}

// NewSimpleExecutor creates a new simple executor.
func NewSimpleExecutor() Executor {
	return &simpleExecutor{}
}

type simpleExecutor struct{}

func (e *simpleExecutor) Execute(strategy strategy.Strategy, wallet wallet.Wallet) error {
	var err error
	for i := 0; i < 3; i++ {
		err = strategy.Execute(wallet)
		if err == nil {
			return nil
		}
		time.Sleep(2 * time.Second)
	}
	return err
}