package strategy

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sheawinkler/farmer-shea/solana"
	"github.com/sheawinkler/farmer-shea/wallet"
)

const (
	MarinadeFinanceProgramID = "MarBmsSgKXdrN1egZf5sqe1TMai9K1rChYNDJgjq7aD"
)

// MarinadeStakingStrategy is a strategy for staking SOL on Marinade Finance.	



type MarinadeStakingStrategy struct {
	solanaClient *solana.Client
	amount       uint64
}

// NewMarinadeStakingStrategy creates a new MarinadeStakingStrategy.
func NewMarinadeStakingStrategy(solanaClient *solana.Client, amount uint64) *MarinadeStakingStrategy {
	return &MarinadeStakingStrategy{
		solanaClient: solanaClient,
		amount:       amount,
	}
}

func (s *MarinadeStakingStrategy) Execute(w wallet.Wallet) error {
	fmt.Printf("Executing Marinade staking strategy with amount %d...\n", s.amount)

	programID, err := solana.PublicKeyFromBase58(MarinadeFinanceProgramID)
	if err != nil {
		return err
	}

	// TODO: Construct the stake instruction based on Marinade's ABI
	// This will involve finding the correct instruction data and accounts.
	ix := solana.NewInstruction(
		programID,
		[]*solana.AccountMeta{
			// Add necessary accounts here
		},
		[]byte{}, // Add instruction data here
	)

	blockhash, err := s.solanaClient.GetLatestBlockhash()
	if err != nil {
		return err
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{ix},
		*blockhash,
		w.PublicKey(),
	)
	if err != nil {
		return err
	}

	if err := w.SignTransaction(tx); err != nil {
		return err
	}

	sig, err := s.solanaClient.RPC().SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}

	return s.solanaClient.RPC().ConfirmTransaction(context.Background(), sig, rpc.CommitmentFinalized)
}