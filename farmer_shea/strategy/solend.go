package strategy

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	bin "github.com/gagliardetto/binary"
	associated_token_account "github.com/gagliardetto/solana-go/programs/associated-token-account"

	"github.com/sheawinkler/farmer-shea/wallet"
	"github.com/sheawinkler/farmer-shea/solana"
	"github.com/sheawinkler/farmer-shea/oracle"
)

const (
	solendProgramID = "So1endDq2YkqhipRh3WViPa8hdiSpxWy6z3Z6tMCpAo"
)

// Solend is a farming strategy for the Solend protocol.	ype Solend struct {
	solanaClient *solana.Client
	oracle       oracle.Oracle
	amount       uint64
	tokenMint    solana.PublicKey
}

// NewSolend creates a new Solend strategy.
func NewSolend(solanaClient *solana.Client, oracle oracle.Oracle, amount uint64, tokenMint solana.PublicKey) *Solend {
	return &Solend{
		solanaClient: solanaClient,
		oracle:       oracle,
		amount:       amount,
		tokenMint:    tokenMint,
	}
}

func (s *Solend) Execute(w wallet.Wallet) error {
	tx, err := s.deposit(s.solanaClient, w, s.amount, s.tokenMint)
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

func (s *Solend) deposit(client *solana.Client, w wallet.Wallet, amount uint64, tokenMint solana.PublicKey) (*solana.Transaction, error) {
	programID, err := solana.PublicKeyFromBase58(solendProgramID)
	if err != nil {
		return nil, err
	}

	reservePubkey, reserve, err := s.findReserveAccount(client, tokenMint)
	if err != nil {
		return nil, err
	}

	userTokenAccount, err := client.GetOrCreateAssociatedTokenAccount(w, tokenMint)
	if err != nil {
		return nil, err
	}

	userCollateralAccount, err := client.GetOrCreateAssociatedTokenAccount(w, reserve.Collateral.MintPubkey)
	if err != nil {
		return nil, err
	}

	lendingMarketAuthority, _, err := solana.FindProgramAddress(
		[][]byte{reserve.LendingMarket.Bytes()},
		programID,
	)
	if err != nil {
		return nil, err
	}

	ix := solana.NewInstruction(
		programID,
		[]*solana.AccountMeta{
			{PublicKey: userTokenAccount, IsSigner: false, IsWritable: true},
			{PublicKey: userCollateralAccount, IsSigner: false, IsWritable: true},
			{PublicKey: *reservePubkey, IsSigner: false, IsWritable: true},
			{PublicKey: reserve.Liquidity.SupplyPubkey, IsSigner: false, IsWritable: true},
			{PublicKey: reserve.Collateral.MintPubkey, IsSigner: false, IsWritable: true},
			{PublicKey: reserve.LendingMarket, IsSigner: false, IsWritable: false},
			{PublicKey: lendingMarketAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: w.PublicKey(), IsSigner: true, IsWritable: false},
			{PublicKey: solana.TokenProgramID, IsSigner: false, IsWritable: false},
		},
		append([]byte{1}, new(bin.Buffer).WriteUint64(amount, bin.LE).Bytes()...),
	)

	blockhash, err := client.GetLatestBlockhash()
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{ix},
		*blockhash,
		w.PublicKey(),
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *Solend) withdraw(client *solana.Client, w wallet.Wallet, amount uint64, tokenMint solana.PublicKey) (*solana.Transaction, error) {
	programID, err := solana.PublicKeyFromBase58(solendProgramID)
	if err != nil {
		return nil, err
	}

	reservePubkey, reserve, err := s.findReserveAccount(client, tokenMint)
	if err != nil {
		return nil, err
	}

	userTokenAccount, err := client.GetOrCreateAssociatedTokenAccount(w, tokenMint)
	if err != nil {
		return nil, err
	}

	userCollateralAccount, err := client.GetOrCreateAssociatedTokenAccount(w, reserve.Collateral.MintPubkey)
	if err != nil {
		return nil, err
	}

	lendingMarketAuthority, _, err := solana.FindProgramAddress(
		[][]byte{reserve.LendingMarket.Bytes()},
		programID,
	)
	if err != nil {
		return nil, err
	}

	ix := solana.NewInstruction(
		programID,
		[]*solana.AccountMeta{
			{PublicKey: userCollateralAccount, IsSigner: false, IsWritable: true},
			{PublicKey: userTokenAccount, IsSigner: false, IsWritable: true},
			{PublicKey: *reservePubkey, IsSigner: false, IsWritable: true},
			{PublicKey: reserve.Collateral.SupplyPubkey, IsSigner: false, IsWritable: true},
			{PublicKey: reserve.LendingMarket, IsSigner: false, IsWritable: false},
			{PublicKey: lendingMarketAuthority, IsSigner: false, IsWritable: false},
			{PublicKey: w.PublicKey(), IsSigner: true, IsWritable: false},
			{PublicKey: solana.TokenProgramID, IsSigner: false, IsWritable: false},
		},
		append([]byte{2}, new(bin.Buffer).WriteUint64(amount, bin.LE).Bytes()...),
	)

	blockhash, err := client.GetLatestBlockhash()
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{ix},
		*blockhash,
		w.PublicKey(),
	)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *Solend) findReserveAccount(client *solana.Client, tokenMint solana.PublicKey) (*solana.PublicKey, *Reserve, error) {
	programID, err := solana.PublicKeyFromBase58(solendProgramID)
	if err != nil {
		return nil, nil, err
	}

	accounts, err := client.GetProgramAccounts(programID.String())
	if err != nil {
		return nil, nil, err
	}

	for _, account := range accounts {
		var reserve Reserve
		if err := bin.NewBinDecoder(account.Account.Data).Decode(&reserve); err != nil {
			continue
		}

		if reserve.Liquidity.MintPubkey == tokenMint {
			return &account.Pubkey, &reserve, nil
		}
	}

	return nil, nil, fmt.Errorf("reserve not found for token mint %s", tokenMint)
}

// Reserve is the structure of a Solend reserve account.	ype Reserve struct {
	Version               uint8
	LastUpdate            LastUpdate
	LendingMarket         solana.PublicKey
	Liquidity             ReserveLiquidity
	Collateral            ReserveCollateral
	Config                ReserveConfig
	Padding               [248]byte
}

// LastUpdate is the structure of the LastUpdate field in a Reserve account.	ype LastUpdate struct {
	Slot    uint64
	Stale   bool
	Padding [7]byte
}

// ReserveLiquidity is the structure of the Liquidity field in a Reserve account.	ype ReserveLiquidity struct {
	MintPubkey            solana.PublicKey
	MintDecimals          uint8
	SupplyPubkey          solana.PublicKey
	PythOraclePubkey      solana.PublicKey
	SwitchboardOraclePubkey solana.PublicKey
	AvailableAmount       uint64
	BorrowedAmountWads    solana.U128
	CumulativeBorrowRateWads solana.U128
	MarketPrice           solana.U128
}

// ReserveCollateral is the structure of the Collateral field in a Reserve account.	ype ReserveCollateral struct {
	MintPubkey           solana.PublicKey
	MintTotalSupply      uint64
	SupplyPubkey         solana.PublicKey
}

// ReserveConfig is the structure of the Config field in a Reserve account.	ype ReserveConfig struct {
	OptimalUtilizationRate  uint8
	LoanToValueRatio        uint8
	LiquidationBonus        uint8
	LiquidationThreshold    uint8
	MinBorrowRate           uint8
	OptimalBorrowRate       uint8
	MaxBorrowRate           uint8
	Fees                    ReserveFees
	DepositLimit            uint64
	BorrowLimit             uint64
	FeeReceiver             solana.PublicKey
	ProtocolLiquidationFee  uint8
	ProtocolTakeRate        uint8
}

// ReserveFees is the structure of the Fees field in a ReserveConfig.	ype ReserveFees struct {
	BorrowFeeWad        uint64
	FlashLoanFeeWad     uint64
	HostFeePercentage   uint8
	Padding             [7]byte
}
