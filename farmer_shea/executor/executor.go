package executor

import (
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
	return strategy.Execute(wallet)
}
