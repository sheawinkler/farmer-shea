package wallet

import (
	"encoding/json"
	"os"

	"github.com/gagliardetto/solana-go"
)

// Wallet represents a Solana wallet.
type Wallet struct {
	PrivateKey solana.PrivateKey `json:"privateKey"`
	PublicKey  solana.PublicKey  `json:"publicKey"`
}

// NewWallet creates a new Solana wallet.
func NewWallet() (*Wallet, error) {
	privateKey, err := solana.NewRandomPrivateKey()
	if err != nil {
		return nil, err
	}

	return &Wallet{
		PrivateKey: privateKey,
		PublicKey:  privateKey.PublicKey(),
	}, nil
}

// Save saves the wallet to a file.
func (w *Wallet) Save(path string) error {
	data, err := json.Marshal(w)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

// Load loads a wallet from a file.
func Load(path string) (*Wallet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var wallet Wallet
	err = json.Unmarshal(data, &wallet)
	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
