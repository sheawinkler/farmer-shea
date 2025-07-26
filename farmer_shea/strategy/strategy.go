package strategy

import (
	"crypto/ecdsa"

	"github.com/sheawinkler/farmer-shea/wallet"
)

// Strategy defines the interface for all trading strategies.
type Strategy interface {
	Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error
	Name() string
}
