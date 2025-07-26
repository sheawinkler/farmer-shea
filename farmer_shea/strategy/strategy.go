package strategy

import (
	"github.com/sheawinkler/farmer-shea/wallet"
)

// Strategy defines the interface for all trading strategies.
type Strategy interface {
	Execute(w wallet.Wallet) error
	Name() string
}
