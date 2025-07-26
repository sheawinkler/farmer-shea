package solana

import (
	"context"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/sheawinkler/farmer-shea/wallet"
)

// Client is a Solana client.
type Client struct {
	*rpc.Client
}

// NewClient creates a new Solana client.
func NewClient(rpcEndpoint string) (*Client, error) {
	client := rpc.New(rpcEndpoint)
	return &Client{client}, nil
}

// GetLatestBlockHeight gets the latest block height of the Solana blockchain.
func (c *Client) GetLatestBlockHeight() (uint64, error) {
	return c.GetSlot(context.Background(), rpc.CommitmentFinalized)
}

// GetOrCreateAssociatedTokenAccount gets or creates an associated token account for the given wallet and mint.
func (c *Client) GetOrCreateAssociatedTokenAccount(w wallet.Wallet, mint solana.PublicKey) (solana.PublicKey, error) {
	ata, _, err := solana.FindAssociatedTokenAddress(w.PublicKey(), mint)
	if err != nil {
		return solana.PublicKey{}, err
	}

	_, err = c.GetAccountInfo(context.Background(), ata)
	if err == nil {
		return ata, nil // Account already exists
	}

	ix, err := associated_token_account.NewCreateInstruction(w.PublicKey(), w.PublicKey(), mint).Validate()
	if err != nil {
		return solana.PublicKey{}, err
	}

	blockhash, err := c.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.PublicKey{}, err
	}

	tx, err := solana.NewTransaction([]solana.Instruction{ix}, blockhash.Value.Blockhash, w.PublicKey())
	if err != nil {
		return solana.PublicKey{}, err
	}

	if err := w.SignTransaction(tx); err != nil {
		return solana.PublicKey{}, err
	}

	sig, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		return solana.PublicKey{}, err
	}

	return ata, c.ConfirmTransaction(context.Background(), sig, rpc.CommitmentFinalized)
}

// GetProgramAccounts gets all accounts owned by a program.
func (c *Client) GetProgramAccounts(programID string) ([]*rpc.GetProgramAccountsResult, error) {
	return c.GetProgramAccounts(context.Background(), solana.MustPublicKeyFromBase58(programID), nil)
}
