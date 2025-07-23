package execution

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/gagliardetto/solana-go/programs/system"
)

// SendSol sends SOL to a given address.
func SendSol(client *rpc.Client, sender solana.PrivateKey, recipient solana.PublicKey, amount uint64) (solana.Signature, error) {
	// Get the recent blockhash
	recentBlockhash, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, err
	}

	// Create the transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amount,
				sender.PublicKey(),
				recipient,
			).Build(),
		},
		recentBlockhash.Value.Blockhash,
		solana.TransactionPayer(sender.PublicKey()),
	)
	if err != nil {
		return solana.Signature{}, err
	}

	// Sign the transaction
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if sender.PublicKey().Equals(key) {
			return &sender
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, fmt.Errorf("unable to sign transaction: %w", err)
	}

	// Send the transaction
	sig, err := client.SendTransaction(context.Background(), tx)
	if err != nil {
		return solana.Signature{}, err
	}

	return sig, nil
}
