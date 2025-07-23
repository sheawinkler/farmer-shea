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

// AdvancedExecutor defines an executor that can handle multiple steps.	ype AdvancedExecutor interface {
	ExecuteSteps(steps []func() error) error
}

// NewAdvancedExecutor creates a new advanced executor.
func NewAdvancedExecutor() AdvancedExecutor {
	return &advancedExecutor{}
}

type advancedExecutor struct{}

func (e *advancedExecutor) ExecuteSteps(steps []func() error) error {
	for _, step := range steps {
		var err error
		for i := 0; i < 3; i++ {
			err = step()
			if err == nil {
				break
			}
			time.Sleep(2 * time.Second)
		}
		if err != nil {
			return err
		}
	}
	return nil
}
