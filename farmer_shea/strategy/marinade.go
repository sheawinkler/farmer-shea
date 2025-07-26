package strategy

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sheawinkler/farmer-shea/solana"
	"github.com/sheawinkler/farmer-shea/wallet"
)

const (
	MarinadeFinanceProgramID = "MarBmsSgKXdrN1egZf5sqe1TMai9K1rChYNDJgjq7aD"
	mSOLMintAddress          = "mSoLzYCxHdYgdzU16g5QSh3i5K3z3KZK7ytfqcJm7So"
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

func (s *MarinadeStakingStrategy) Name() string {
	return "MarinadeStaking"
}

func (s *MarinadeStakingStrategy) Execute(w wallet.Wallet, privateKey *ecdsa.PrivateKey) error {
	fmt.Printf("Executing Marinade staking strategy with amount %d...\n", s.amount)

	programID, err := solana.PublicKeyFromBase58(MarinadeFinanceProgramID)
	if err != nil {
		return err
	}

	mSOLMint, err := solana.PublicKeyFromBase58(mSOLMintAddress)
	if err != nil {
		return err
	}

	state, err := s.getMarinadeState(programID)
	if err != nil {
		return err
	}

	mSOLTokenAccount, err := s.solanaClient.GetOrCreateAssociatedTokenAccount(w, mSOLMint)
	if err != nil {
		return err
	}

	ix := s.createDepositInstruction(w.PublicKey(), programID, state, mSOLTokenAccount, s.amount)

	blockhash, err := s.solanaClient.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return err
	}

	tx, err := solana.NewTransaction([]solana.Instruction{ix}, blockhash.Value.Blockhash, w.PublicKey())
	if err != nil {
		return err
	}

	if err := w.SignTransaction(tx); err != nil {
		return err
	}

	sig, err := s.solanaClient.SendTransaction(context.Background(), tx)
	if err != nil {
		return err
	}

	return s.solanaClient.ConfirmTransaction(context.Background(), sig, rpc.CommitmentFinalized)
}

func (s *MarinadeStakingStrategy) getMarinadeState(programID solana.PublicKey) (*solana.PublicKey, error) {
	// This is a simplified way to get the state address. A more robust implementation
	// would involve querying for the account with the correct data.
	state, _, err := solana.FindProgramAddress([][]byte{[]byte("state")}, programID)
	return &state, err
}

func (s *MarinadeStakingStrategy) createDepositInstruction(user solana.PublicKey, programID, state, mSOLTokenAccount solana.PublicKey, amount uint64) solana.Instruction {
	return system.NewTransferInstruction(amount, user, state).Build()
}
