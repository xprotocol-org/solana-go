// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package memo_transfer

import (
	"errors"
	common "github.com/gagliardetto/solana-go/programs/common"
	binary "github.com/gagliardetto/binary"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Enable Instruction
type Enable struct {
	// [0] = [WRITE] account `The account to update.`
	// [1] = [SIGNER] accountOwner `The account's owner. or The account's multisignature owner.,..2+M `[signer]` M signer accounts.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewEnableInstructionBuilder creates a new `Enable` instruction builder.
func NewEnableInstructionBuilder() *Enable {
	return &Enable{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewEnableInstruction
//
// Parameters:
//   account: The account to update.
/*
  accountOwner: The account's owner. or The account's multisignature owner.
  ..2+M `[signer]` M signer accounts.
*/
//
func NewEnableInstruction(
	account common.PublicKey,
	accountOwner common.PublicKey,
) *Enable {
	return NewEnableInstructionBuilder().
		SetAccountAccount(account).
		SetAccountOwnerAccount(accountOwner)
}

// SetAccountAccount sets the "account" parameter.
// The account to update.
func (obj *Enable) SetAccountAccount(account common.PublicKey) *Enable {
	obj.AccountMetaSlice[0] = common.Meta(account).WRITE()
	return obj
}

// GetAccountAccount gets the "account" parameter.
// The account to update.
func (obj *Enable) GetAccountAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetAccountOwnerAccount sets the "accountOwner" parameter.
// The account's owner. or The account's multisignature owner.
// ..2+M `[signer]` M signer accounts.
func (obj *Enable) SetAccountOwnerAccount(accountOwner common.PublicKey, multiSigners ...common.PublicKey) *Enable {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(accountOwner)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(accountOwner).SIGNER()
	}
	return obj
}

// GetAccountOwnerAccount gets the "accountOwner" parameter.
// The account's owner. or The account's multisignature owner.
// ..2+M `[signer]` M signer accounts.
func (obj *Enable) GetAccountOwnerAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *Enable) SetProgramId(programId *common.PublicKey) *Enable {
	obj._programId = programId
	return obj
}

func (obj *Enable) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Enable}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Enable) Validate() error {

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Enable] accounts.account is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[Enable] accounts.accountOwner is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Enable) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Enable) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	return nil
}

func (obj *Enable) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	return nil
}

func (obj *Enable) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Enable")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch treeout.Branches) {})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("     account", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("accountOwner", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}

// Disable Instruction
type Disable struct {
	// [0] = [WRITE] account `The account to update.`
	// [1] = [SIGNER] accountOwner `The account's owner. or The account's multisignature owner.,..2+M `[signer]` M signer accounts.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewDisableInstructionBuilder creates a new `Disable` instruction builder.
func NewDisableInstructionBuilder() *Disable {
	return &Disable{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewDisableInstruction
//
// Parameters:
//   account: The account to update.
/*
  accountOwner: The account's owner. or The account's multisignature owner.
  ..2+M `[signer]` M signer accounts.
*/
//
func NewDisableInstruction(
	account common.PublicKey,
	accountOwner common.PublicKey,
) *Disable {
	return NewDisableInstructionBuilder().
		SetAccountAccount(account).
		SetAccountOwnerAccount(accountOwner)
}

// SetAccountAccount sets the "account" parameter.
// The account to update.
func (obj *Disable) SetAccountAccount(account common.PublicKey) *Disable {
	obj.AccountMetaSlice[0] = common.Meta(account).WRITE()
	return obj
}

// GetAccountAccount gets the "account" parameter.
// The account to update.
func (obj *Disable) GetAccountAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetAccountOwnerAccount sets the "accountOwner" parameter.
// The account's owner. or The account's multisignature owner.
// ..2+M `[signer]` M signer accounts.
func (obj *Disable) SetAccountOwnerAccount(accountOwner common.PublicKey, multiSigners ...common.PublicKey) *Disable {
	if len(multiSigners) > 0 {
		obj.AccountMetaSlice[1] = common.Meta(accountOwner)
		for _, value := range multiSigners {
			obj.AccountMetaSlice.Append(common.NewAccountMeta(value, false, true))
		}
	} else {
		obj.AccountMetaSlice[1] = common.Meta(accountOwner).SIGNER()
	}
	return obj
}

// GetAccountOwnerAccount gets the "accountOwner" parameter.
// The account's owner. or The account's multisignature owner.
// ..2+M `[signer]` M signer accounts.
func (obj *Disable) GetAccountOwnerAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *Disable) SetProgramId(programId *common.PublicKey) *Disable {
	obj._programId = programId
	return obj
}

func (obj *Disable) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Disable}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Disable) Validate() error {

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Disable] accounts.account is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[Disable] accounts.accountOwner is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Disable) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Disable) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	return nil
}

func (obj *Disable) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	return nil
}

func (obj *Disable) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Disable")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch treeout.Branches) {})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("     account", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("accountOwner", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}
