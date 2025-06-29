// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package transfer_fee

import (
	"errors"
	common "github.com/gagliardetto/solana-go/programs/common"
	binary "github.com/gagliardetto/binary"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// InitializeTransferFeeConfig Instruction
type InitializeTransferFeeConfig struct {
	// Pubkey that may update the fees
	TransferFeeConfigAuthority *common.PublicKey `bin:"optional"`
	// Withdraw instructions must be signed by this key
	WithdrawWithheldAuthority *common.PublicKey `bin:"optional"`
	// Amount of transfer collected as fees, expressed as basis points of,the transfer amount
	TransferFeeBasisPoints *uint16
	// Maximum fee assessed on transfers
	MaximumFee *uint64
	// [0] = [WRITE] mint `The mint to initialize`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewInitializeTransferFeeConfigInstructionBuilder creates a new `InitializeTransferFeeConfig` instruction builder.
func NewInitializeTransferFeeConfigInstructionBuilder() *InitializeTransferFeeConfig {
	return &InitializeTransferFeeConfig{
		AccountMetaSlice: make(common.AccountMetaSlice, 1),
	}
}

// NewInitializeTransferFeeConfigInstruction
//
// Parameters:
//   transferFeeConfigAuthority: Pubkey that may update the fees
//   withdrawWithheldAuthority: Withdraw instructions must be signed by this key
/*
  transferFeeBasisPoints: Amount of transfer collected as fees, expressed as basis points of
  the transfer amount
*/
//   maximumFee: Maximum fee assessed on transfers
//   mint: The mint to initialize
//
func NewInitializeTransferFeeConfigInstruction(
	// optional,
	transferFeeConfigAuthority *common.PublicKey,
	// optional,
	withdrawWithheldAuthority *common.PublicKey,
	transferFeeBasisPoints uint16,
	maximumFee uint64,
	mint common.PublicKey,
) *InitializeTransferFeeConfig {
	return NewInitializeTransferFeeConfigInstructionBuilder().
		SetTransferFeeConfigAuthority(transferFeeConfigAuthority).
		SetWithdrawWithheldAuthority(withdrawWithheldAuthority).
		SetTransferFeeBasisPoints(transferFeeBasisPoints).
		SetMaximumFee(maximumFee).
		SetMintAccount(mint)
}

// SetTransferFeeConfigAuthority sets the "transferFeeConfigAuthority" parameter.
func (obj *InitializeTransferFeeConfig) SetTransferFeeConfigAuthority(transferFeeConfigAuthority *common.PublicKey) *InitializeTransferFeeConfig {
	obj.TransferFeeConfigAuthority = transferFeeConfigAuthority
	return obj
}

// SetWithdrawWithheldAuthority sets the "withdrawWithheldAuthority" parameter.
func (obj *InitializeTransferFeeConfig) SetWithdrawWithheldAuthority(withdrawWithheldAuthority *common.PublicKey) *InitializeTransferFeeConfig {
	obj.WithdrawWithheldAuthority = withdrawWithheldAuthority
	return obj
}

// SetTransferFeeBasisPoints sets the "transferFeeBasisPoints" parameter.
func (obj *InitializeTransferFeeConfig) SetTransferFeeBasisPoints(transferFeeBasisPoints uint16) *InitializeTransferFeeConfig {
	obj.TransferFeeBasisPoints = &transferFeeBasisPoints
	return obj
}

// SetMaximumFee sets the "maximumFee" parameter.
func (obj *InitializeTransferFeeConfig) SetMaximumFee(maximumFee uint64) *InitializeTransferFeeConfig {
	obj.MaximumFee = &maximumFee
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint to initialize
func (obj *InitializeTransferFeeConfig) SetMintAccount(mint common.PublicKey, multiSigners ...common.PublicKey) *InitializeTransferFeeConfig {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[0] = common.Meta(mint)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	}
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint to initialize
func (obj *InitializeTransferFeeConfig) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

func (obj *InitializeTransferFeeConfig) SetProgramId(programId *common.PublicKey) *InitializeTransferFeeConfig {
	obj._programId = programId
	return obj
}

func (obj *InitializeTransferFeeConfig) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_InitializeTransferFeeConfig}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *InitializeTransferFeeConfig) Validate() error {
	if obj.TransferFeeBasisPoints == nil {
		return errors.New("[InitializeTransferFeeConfig] transferFeeBasisPoints param is not set")
	}
	if obj.MaximumFee == nil {
		return errors.New("[InitializeTransferFeeConfig] maximumFee param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[InitializeTransferFeeConfig] accounts.mint is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *InitializeTransferFeeConfig) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *InitializeTransferFeeConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.WriteBool(obj.TransferFeeConfigAuthority != nil); err != nil {
		return err
	}
	if obj.TransferFeeConfigAuthority != nil {
		if err = encoder.Encode(obj.TransferFeeConfigAuthority); err != nil {
			return err
		}
	}
	if err = encoder.WriteBool(obj.WithdrawWithheldAuthority != nil); err != nil {
		return err
	}
	if obj.WithdrawWithheldAuthority != nil {
		if err = encoder.Encode(obj.WithdrawWithheldAuthority); err != nil {
			return err
		}
	}
	if err = encoder.Encode(&obj.TransferFeeBasisPoints); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.MaximumFee); err != nil {
		return err
	}
	return nil
}

func (obj *InitializeTransferFeeConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if ok, err := decoder.ReadBool(); err != nil {
		return err
	} else if ok {
		if err = decoder.Decode(&obj.TransferFeeConfigAuthority); err != nil {
			return err
		}
	}
	if ok, err := decoder.ReadBool(); err != nil {
		return err
	} else if ok {
		if err = decoder.Decode(&obj.WithdrawWithheldAuthority); err != nil {
			return err
		}
	}
	if err = decoder.Decode(&obj.TransferFeeBasisPoints); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.MaximumFee); err != nil {
		return err
	}
	return nil
}

func (obj *InitializeTransferFeeConfig) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("InitializeTransferFeeConfig")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=4]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("TransferFeeConfigAuthority (OPT)", obj.TransferFeeConfigAuthority))
						paramsBranch.Child(format.Param(" WithdrawWithheldAuthority (OPT)", obj.WithdrawWithheldAuthority))
						paramsBranch.Child(format.Param("    TransferFeeBasisPoints", *obj.TransferFeeBasisPoints))
						paramsBranch.Child(format.Param("                MaximumFee", *obj.MaximumFee))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("mint", obj.AccountMetaSlice.Get(0)))
					})
				})
		})
}

// TransferCheckedWithFee Instruction
type TransferCheckedWithFee struct {
	// The amount of tokens to transfer.
	Amount *uint64
	// Expected number of base 10 digits to the right of the decimal place.
	Decimals *uint8
	// Expected fee assessed on this transfer, calculated off-chain based,on the transfer_fee_basis_points and maximum_fee of the,mint.
	Fee *uint64
	// [0] = [WRITE] source `The source account. Must include the `TransferFeeAmount` extension.`
	// [1] = [] mint `The token mint. Must include the `TransferFeeConfig` extension.`
	// [2] = [] destination `The destination account. Must include the `TransferFeeAmount` extension.`
	// [3] = [SIGNER] authority `The source account's owner/delegate.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewTransferCheckedWithFeeInstructionBuilder creates a new `TransferCheckedWithFee` instruction builder.
func NewTransferCheckedWithFeeInstructionBuilder() *TransferCheckedWithFee {
	return &TransferCheckedWithFee{
		AccountMetaSlice: make(common.AccountMetaSlice, 4),
	}
}

// NewTransferCheckedWithFeeInstruction
//
// Parameters:
//   amount: The amount of tokens to transfer.
//   decimals: Expected number of base 10 digits to the right of the decimal place.
/*
  fee: Expected fee assessed on this transfer, calculated off-chain based
  on the transfer_fee_basis_points and maximum_fee of the
  mint.
*/
//   source: The source account. Must include the `TransferFeeAmount` extension.
//   mint: The token mint. Must include the `TransferFeeConfig` extension.
//   destination: The destination account. Must include the `TransferFeeAmount` extension.
//   authority: The source account's owner/delegate.
//
func NewTransferCheckedWithFeeInstruction(
	amount uint64,
	decimals uint8,
	fee uint64,
	source common.PublicKey,
	mint common.PublicKey,
	destination common.PublicKey,
	authority common.PublicKey,
) *TransferCheckedWithFee {
	return NewTransferCheckedWithFeeInstructionBuilder().
		SetAmount(amount).
		SetDecimals(decimals).
		SetFee(fee).
		SetSourceAccount(source).
		SetMintAccount(mint).
		SetDestinationAccount(destination).
		SetAuthorityAccount(authority)
}

// SetAmount sets the "amount" parameter.
func (obj *TransferCheckedWithFee) SetAmount(amount uint64) *TransferCheckedWithFee {
	obj.Amount = &amount
	return obj
}

// SetDecimals sets the "decimals" parameter.
func (obj *TransferCheckedWithFee) SetDecimals(decimals uint8) *TransferCheckedWithFee {
	obj.Decimals = &decimals
	return obj
}

// SetFee sets the "fee" parameter.
func (obj *TransferCheckedWithFee) SetFee(fee uint64) *TransferCheckedWithFee {
	obj.Fee = &fee
	return obj
}

// SetSourceAccount sets the "source" parameter.
// The source account. Must include the `TransferFeeAmount` extension.
func (obj *TransferCheckedWithFee) SetSourceAccount(source common.PublicKey) *TransferCheckedWithFee {
	obj.AccountMetaSlice[0] = common.Meta(source).WRITE()
	return obj
}

// GetSourceAccount gets the "source" parameter.
// The source account. Must include the `TransferFeeAmount` extension.
func (obj *TransferCheckedWithFee) GetSourceAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetMintAccount sets the "mint" parameter.
// The token mint. Must include the `TransferFeeConfig` extension.
func (obj *TransferCheckedWithFee) SetMintAccount(mint common.PublicKey) *TransferCheckedWithFee {
	obj.AccountMetaSlice[1] = common.Meta(mint)
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The token mint. Must include the `TransferFeeConfig` extension.
func (obj *TransferCheckedWithFee) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

// SetDestinationAccount sets the "destination" parameter.
// The destination account. Must include the `TransferFeeAmount` extension.
func (obj *TransferCheckedWithFee) SetDestinationAccount(destination common.PublicKey) *TransferCheckedWithFee {
	obj.AccountMetaSlice[2] = common.Meta(destination)
	return obj
}

// GetDestinationAccount gets the "destination" parameter.
// The destination account. Must include the `TransferFeeAmount` extension.
func (obj *TransferCheckedWithFee) GetDestinationAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(2)
}

// SetAuthorityAccount sets the "authority" parameter.
// The source account's owner/delegate.
func (obj *TransferCheckedWithFee) SetAuthorityAccount(authority common.PublicKey, multiSigners ...common.PublicKey) *TransferCheckedWithFee {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[3] = common.Meta(authority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[3] = common.Meta(authority).SIGNER()
	}
	return obj
}

// GetAuthorityAccount gets the "authority" parameter.
// The source account's owner/delegate.
func (obj *TransferCheckedWithFee) GetAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(3)
}

func (obj *TransferCheckedWithFee) SetProgramId(programId *common.PublicKey) *TransferCheckedWithFee {
	obj._programId = programId
	return obj
}

func (obj *TransferCheckedWithFee) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_TransferCheckedWithFee}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *TransferCheckedWithFee) Validate() error {
	if obj.Amount == nil {
		return errors.New("[TransferCheckedWithFee] amount param is not set")
	}
	if obj.Decimals == nil {
		return errors.New("[TransferCheckedWithFee] decimals param is not set")
	}
	if obj.Fee == nil {
		return errors.New("[TransferCheckedWithFee] fee param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[TransferCheckedWithFee] accounts.source is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[TransferCheckedWithFee] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[2] == nil {
		return errors.New("[TransferCheckedWithFee] accounts.destination is not set")
	}
	if obj.AccountMetaSlice[3] == nil {
		return errors.New("[TransferCheckedWithFee] accounts.authority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *TransferCheckedWithFee) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *TransferCheckedWithFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Amount); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Decimals); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Fee); err != nil {
		return err
	}
	return nil
}

func (obj *TransferCheckedWithFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Amount); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Decimals); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Fee); err != nil {
		return err
	}
	return nil
}

func (obj *TransferCheckedWithFee) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("TransferCheckedWithFee")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=3]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("  Amount", *obj.Amount))
						paramsBranch.Child(format.Param("Decimals", *obj.Decimals))
						paramsBranch.Child(format.Param("     Fee", *obj.Fee))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=4]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("     source", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("       mint", obj.AccountMetaSlice.Get(1)))
						accountsBranch.Child(common.FormatMeta("destination", obj.AccountMetaSlice.Get(2)))
						accountsBranch.Child(common.FormatMeta("  authority", obj.AccountMetaSlice.Get(3)))
					})
				})
		})
}

// WithdrawWithheldTokensFromMint Instruction
type WithdrawWithheldTokensFromMint struct {
	// [0] = [WRITE] mint `The token mint. Must include the `TransferFeeConfig` extension.`
	// [1] = [WRITE] feeReceiver `The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.`
	// [2] = [SIGNER] authority `The mint's `withdraw_withheld_authority`.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewWithdrawWithheldTokensFromMintInstructionBuilder creates a new `WithdrawWithheldTokensFromMint` instruction builder.
func NewWithdrawWithheldTokensFromMintInstructionBuilder() *WithdrawWithheldTokensFromMint {
	return &WithdrawWithheldTokensFromMint{
		AccountMetaSlice: make(common.AccountMetaSlice, 3),
	}
}

// NewWithdrawWithheldTokensFromMintInstruction
//
// Parameters:
//
//	mint: The token mint. Must include the `TransferFeeConfig` extension.
//	feeReceiver: The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.
//	authority: The mint's `withdraw_withheld_authority`.
func NewWithdrawWithheldTokensFromMintInstruction(
	mint common.PublicKey,
	feeReceiver common.PublicKey,
	authority common.PublicKey,
) *WithdrawWithheldTokensFromMint {
	return NewWithdrawWithheldTokensFromMintInstructionBuilder().
		SetMintAccount(mint).
		SetFeeReceiverAccount(feeReceiver).
		SetAuthorityAccount(authority)
}

// SetMintAccount sets the "mint" parameter.
// The token mint. Must include the `TransferFeeConfig` extension.
func (obj *WithdrawWithheldTokensFromMint) SetMintAccount(mint common.PublicKey) *WithdrawWithheldTokensFromMint {
	obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The token mint. Must include the `TransferFeeConfig` extension.
func (obj *WithdrawWithheldTokensFromMint) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetFeeReceiverAccount sets the "feeReceiver" parameter.
// The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.
func (obj *WithdrawWithheldTokensFromMint) SetFeeReceiverAccount(feeReceiver common.PublicKey) *WithdrawWithheldTokensFromMint {
	obj.AccountMetaSlice[1] = common.Meta(feeReceiver).WRITE()
	return obj
}

// GetFeeReceiverAccount gets the "feeReceiver" parameter.
// The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.
func (obj *WithdrawWithheldTokensFromMint) GetFeeReceiverAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" parameter.
// The mint's `withdraw_withheld_authority`.
func (obj *WithdrawWithheldTokensFromMint) SetAuthorityAccount(authority common.PublicKey, multiSigners ...common.PublicKey) *WithdrawWithheldTokensFromMint {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[2] = common.Meta(authority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[2] = common.Meta(authority).SIGNER()
	}
	return obj
}

// GetAuthorityAccount gets the "authority" parameter.
// The mint's `withdraw_withheld_authority`.
func (obj *WithdrawWithheldTokensFromMint) GetAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(2)
}

func (obj *WithdrawWithheldTokensFromMint) SetProgramId(programId *common.PublicKey) *WithdrawWithheldTokensFromMint {
	obj._programId = programId
	return obj
}

func (obj *WithdrawWithheldTokensFromMint) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_WithdrawWithheldTokensFromMint}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *WithdrawWithheldTokensFromMint) Validate() error {

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[WithdrawWithheldTokensFromMint] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[WithdrawWithheldTokensFromMint] accounts.feeReceiver is not set")
	}
	if obj.AccountMetaSlice[2] == nil {
		return errors.New("[WithdrawWithheldTokensFromMint] accounts.authority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *WithdrawWithheldTokensFromMint) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *WithdrawWithheldTokensFromMint) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	return nil
}

func (obj *WithdrawWithheldTokensFromMint) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	return nil
}

func (obj *WithdrawWithheldTokensFromMint) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("WithdrawWithheldTokensFromMint")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch treeout.Branches) {})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("       mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("feeReceiver", obj.AccountMetaSlice.Get(1)))
						accountsBranch.Child(common.FormatMeta("  authority", obj.AccountMetaSlice.Get(2)))
					})
				})
		})
}

// WithdrawWithheldTokensFromAccounts Instruction
type WithdrawWithheldTokensFromAccounts struct {
	// Number of token accounts harvested
	NumTokenAccounts *uint8
	// [0] = [] mint `The token mint. Must include the `TransferFeeConfig` extension.`
	// [1] = [WRITE] feeReceiver `The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.`
	// [2] = [SIGNER] authority `The mint's `withdraw_withheld_authority`.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewWithdrawWithheldTokensFromAccountsInstructionBuilder creates a new `WithdrawWithheldTokensFromAccounts` instruction builder.
func NewWithdrawWithheldTokensFromAccountsInstructionBuilder() *WithdrawWithheldTokensFromAccounts {
	return &WithdrawWithheldTokensFromAccounts{
		AccountMetaSlice: make(common.AccountMetaSlice, 3),
	}
}

// NewWithdrawWithheldTokensFromAccountsInstruction
//
// Parameters:
//
//	numTokenAccounts: Number of token accounts harvested
//	mint: The token mint. Must include the `TransferFeeConfig` extension.
//	feeReceiver: The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.
//	authority: The mint's `withdraw_withheld_authority`.
func NewWithdrawWithheldTokensFromAccountsInstruction(
	numTokenAccounts uint8,
	mint common.PublicKey,
	feeReceiver common.PublicKey,
	authority common.PublicKey,
) *WithdrawWithheldTokensFromAccounts {
	return NewWithdrawWithheldTokensFromAccountsInstructionBuilder().
		SetNumTokenAccounts(numTokenAccounts).
		SetMintAccount(mint).
		SetFeeReceiverAccount(feeReceiver).
		SetAuthorityAccount(authority)
}

// SetNumTokenAccounts sets the "numTokenAccounts" parameter.
func (obj *WithdrawWithheldTokensFromAccounts) SetNumTokenAccounts(numTokenAccounts uint8) *WithdrawWithheldTokensFromAccounts {
	obj.NumTokenAccounts = &numTokenAccounts
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The token mint. Must include the `TransferFeeConfig` extension.
func (obj *WithdrawWithheldTokensFromAccounts) SetMintAccount(mint common.PublicKey) *WithdrawWithheldTokensFromAccounts {
	obj.AccountMetaSlice[0] = common.Meta(mint)
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The token mint. Must include the `TransferFeeConfig` extension.
func (obj *WithdrawWithheldTokensFromAccounts) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetFeeReceiverAccount sets the "feeReceiver" parameter.
// The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.
func (obj *WithdrawWithheldTokensFromAccounts) SetFeeReceiverAccount(feeReceiver common.PublicKey) *WithdrawWithheldTokensFromAccounts {
	obj.AccountMetaSlice[1] = common.Meta(feeReceiver).WRITE()
	return obj
}

// GetFeeReceiverAccount gets the "feeReceiver" parameter.
// The fee receiver account. Must include the `TransferFeeAmount` extension associated with the provided mint.
func (obj *WithdrawWithheldTokensFromAccounts) GetFeeReceiverAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

// SetAuthorityAccount sets the "authority" parameter.
// The mint's `withdraw_withheld_authority`.
func (obj *WithdrawWithheldTokensFromAccounts) SetAuthorityAccount(authority common.PublicKey, multiSigners ...common.PublicKey) *WithdrawWithheldTokensFromAccounts {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[2] = common.Meta(authority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[2] = common.Meta(authority).SIGNER()
	}
	return obj
}

// GetAuthorityAccount gets the "authority" parameter.
// The mint's `withdraw_withheld_authority`.
func (obj *WithdrawWithheldTokensFromAccounts) GetAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(2)
}

func (obj *WithdrawWithheldTokensFromAccounts) SetProgramId(programId *common.PublicKey) *WithdrawWithheldTokensFromAccounts {
	obj._programId = programId
	return obj
}

func (obj *WithdrawWithheldTokensFromAccounts) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_WithdrawWithheldTokensFromAccounts}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *WithdrawWithheldTokensFromAccounts) Validate() error {
	if obj.NumTokenAccounts == nil {
		return errors.New("[WithdrawWithheldTokensFromAccounts] numTokenAccounts param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[WithdrawWithheldTokensFromAccounts] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[WithdrawWithheldTokensFromAccounts] accounts.feeReceiver is not set")
	}
	if obj.AccountMetaSlice[2] == nil {
		return errors.New("[WithdrawWithheldTokensFromAccounts] accounts.authority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *WithdrawWithheldTokensFromAccounts) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *WithdrawWithheldTokensFromAccounts) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.NumTokenAccounts); err != nil {
		return err
	}
	return nil
}

func (obj *WithdrawWithheldTokensFromAccounts) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.NumTokenAccounts); err != nil {
		return err
	}
	return nil
}

func (obj *WithdrawWithheldTokensFromAccounts) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("WithdrawWithheldTokensFromAccounts")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("NumTokenAccounts", *obj.NumTokenAccounts))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=3]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("       mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("feeReceiver", obj.AccountMetaSlice.Get(1)))
						accountsBranch.Child(common.FormatMeta("  authority", obj.AccountMetaSlice.Get(2)))
					})
				})
		})
}

// HarvestWithheldTokensToMint Instruction
type HarvestWithheldTokensToMint struct {
	// [0] = [WRITE] mint `The mint.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewHarvestWithheldTokensToMintInstructionBuilder creates a new `HarvestWithheldTokensToMint` instruction builder.
func NewHarvestWithheldTokensToMintInstructionBuilder() *HarvestWithheldTokensToMint {
	return &HarvestWithheldTokensToMint{
		AccountMetaSlice: make(common.AccountMetaSlice, 1),
	}
}

// NewHarvestWithheldTokensToMintInstruction
//
// Parameters:
//
//	mint: The mint.
func NewHarvestWithheldTokensToMintInstruction(
	mint common.PublicKey,
) *HarvestWithheldTokensToMint {
	return NewHarvestWithheldTokensToMintInstructionBuilder().
		SetMintAccount(mint)
}

// SetMintAccount sets the "mint" parameter.
// The mint.
func (obj *HarvestWithheldTokensToMint) SetMintAccount(mint common.PublicKey, multiSigners ...common.PublicKey) *HarvestWithheldTokensToMint {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[0] = common.Meta(mint)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	}
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint.
func (obj *HarvestWithheldTokensToMint) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

func (obj *HarvestWithheldTokensToMint) SetProgramId(programId *common.PublicKey) *HarvestWithheldTokensToMint {
	obj._programId = programId
	return obj
}

func (obj *HarvestWithheldTokensToMint) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_HarvestWithheldTokensToMint}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *HarvestWithheldTokensToMint) Validate() error {

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[HarvestWithheldTokensToMint] accounts.mint is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *HarvestWithheldTokensToMint) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *HarvestWithheldTokensToMint) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	return nil
}

func (obj *HarvestWithheldTokensToMint) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	return nil
}

func (obj *HarvestWithheldTokensToMint) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("HarvestWithheldTokensToMint")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch treeout.Branches) {})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("mint", obj.AccountMetaSlice.Get(0)))
					})
				})
		})
}

// SetTransferFee Instruction
type SetTransferFee struct {
	// Amount of transfer collected as fees, expressed as basis points of,the transfer amount
	TransferFeeBasisPoints *uint16
	// Maximum fee assessed on transfers
	MaximumFee *uint64
	// [0] = [WRITE] mint `The mint.`
	// [1] = [SIGNER] authority `The mint's fee account owner.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewSetTransferFeeInstructionBuilder creates a new `SetTransferFee` instruction builder.
func NewSetTransferFeeInstructionBuilder() *SetTransferFee {
	return &SetTransferFee{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewSetTransferFeeInstruction
//
// Parameters:
/*
  transferFeeBasisPoints: Amount of transfer collected as fees, expressed as basis points of
  the transfer amount
*/
//   maximumFee: Maximum fee assessed on transfers
//   mint: The mint.
//   authority: The mint's fee account owner.
//
func NewSetTransferFeeInstruction(
	transferFeeBasisPoints uint16,
	maximumFee uint64,
	mint common.PublicKey,
	authority common.PublicKey,
) *SetTransferFee {
	return NewSetTransferFeeInstructionBuilder().
		SetTransferFeeBasisPoints(transferFeeBasisPoints).
		SetMaximumFee(maximumFee).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}

// SetTransferFeeBasisPoints sets the "transferFeeBasisPoints" parameter.
func (obj *SetTransferFee) SetTransferFeeBasisPoints(transferFeeBasisPoints uint16) *SetTransferFee {
	obj.TransferFeeBasisPoints = &transferFeeBasisPoints
	return obj
}

// SetMaximumFee sets the "maximumFee" parameter.
func (obj *SetTransferFee) SetMaximumFee(maximumFee uint64) *SetTransferFee {
	obj.MaximumFee = &maximumFee
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint.
func (obj *SetTransferFee) SetMintAccount(mint common.PublicKey) *SetTransferFee {
	obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint.
func (obj *SetTransferFee) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" parameter.
// The mint's fee account owner.
func (obj *SetTransferFee) SetAuthorityAccount(authority common.PublicKey, multiSigners ...common.PublicKey) *SetTransferFee {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(authority)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(authority).SIGNER()
	}
	return obj
}

// GetAuthorityAccount gets the "authority" parameter.
// The mint's fee account owner.
func (obj *SetTransferFee) GetAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *SetTransferFee) SetProgramId(programId *common.PublicKey) *SetTransferFee {
	obj._programId = programId
	return obj
}

func (obj *SetTransferFee) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_SetTransferFee}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *SetTransferFee) Validate() error {
	if obj.TransferFeeBasisPoints == nil {
		return errors.New("[SetTransferFee] transferFeeBasisPoints param is not set")
	}
	if obj.MaximumFee == nil {
		return errors.New("[SetTransferFee] maximumFee param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[SetTransferFee] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[SetTransferFee] accounts.authority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *SetTransferFee) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *SetTransferFee) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.TransferFeeBasisPoints); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.MaximumFee); err != nil {
		return err
	}
	return nil
}

func (obj *SetTransferFee) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.TransferFeeBasisPoints); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.MaximumFee); err != nil {
		return err
	}
	return nil
}

func (obj *SetTransferFee) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("SetTransferFee")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("TransferFeeBasisPoints", *obj.TransferFeeBasisPoints))
						paramsBranch.Child(format.Param("            MaximumFee", *obj.MaximumFee))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("     mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("authority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}
