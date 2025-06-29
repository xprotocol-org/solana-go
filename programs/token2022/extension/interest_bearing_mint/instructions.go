// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package interest_bearing_mint

import (
	"errors"
	common "github.com/gagliardetto/solana-go/programs/common"
	binary "github.com/gagliardetto/binary"
	format "github.com/gagliardetto/solana-go/text/format"
	treeout "github.com/gagliardetto/treeout"
)

// Initialize Instruction
type Initialize struct {
	// The public key for the account that can update the rate
	RateAuthority *common.PublicKey
	// The initial interest rate
	Rate *int16
	// [0] = [WRITE] mint `The mint to initialize`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	return &Initialize{
		AccountMetaSlice: make(common.AccountMetaSlice, 1),
	}
}

// NewInitializeInstruction
//
// Parameters:
//
//	rateAuthority: The public key for the account that can update the rate
//	rate: The initial interest rate
//	mint: The mint to initialize
func NewInitializeInstruction(
	rateAuthority common.PublicKey,
	rate int16,
	mint common.PublicKey,
) *Initialize {
	return NewInitializeInstructionBuilder().
		SetRateAuthority(rateAuthority).
		SetRate(rate).
		SetMintAccount(mint)
}

// SetRateAuthority sets the "rateAuthority" parameter.
func (obj *Initialize) SetRateAuthority(rateAuthority common.PublicKey) *Initialize {
	obj.RateAuthority = &rateAuthority
	return obj
}

// SetRate sets the "rate" parameter.
func (obj *Initialize) SetRate(rate int16) *Initialize {
	obj.Rate = &rate
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint to initialize
func (obj *Initialize) SetMintAccount(mint common.PublicKey, multiSigners ...common.PublicKey) *Initialize {
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
func (obj *Initialize) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

func (obj *Initialize) SetProgramId(programId *common.PublicKey) *Initialize {
	obj._programId = programId
	return obj
}

func (obj *Initialize) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_Initialize}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *Initialize) Validate() error {
	if obj.RateAuthority == nil {
		return errors.New("[Initialize] rateAuthority param is not set")
	}
	if obj.Rate == nil {
		return errors.New("[Initialize] rate param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[Initialize] accounts.mint is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *Initialize) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.RateAuthority); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.Rate); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.RateAuthority); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.Rate); err != nil {
		return err
	}
	return nil
}

func (obj *Initialize) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("Initialize")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("RateAuthority", *obj.RateAuthority))
						paramsBranch.Child(format.Param("         Rate", *obj.Rate))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=1]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("mint", obj.AccountMetaSlice.Get(0)))
					})
				})
		})
}

// UpdateRate Instruction
type UpdateRate struct {
	Rate *int16
	// [0] = [WRITE] mint `The mint`
	// [1] = [SIGNER] authority `The mint rate authority.`
	common.AccountMetaSlice `bin:"-"`
	_programId              *common.PublicKey
}

// NewUpdateRateInstructionBuilder creates a new `UpdateRate` instruction builder.
func NewUpdateRateInstructionBuilder() *UpdateRate {
	return &UpdateRate{
		AccountMetaSlice: make(common.AccountMetaSlice, 2),
	}
}

// NewUpdateRateInstruction
//
// Parameters:
//
//	rate:
//	mint: The mint
//	authority: The mint rate authority.
func NewUpdateRateInstruction(
	rate int16,
	mint common.PublicKey,
	authority common.PublicKey,
) *UpdateRate {
	return NewUpdateRateInstructionBuilder().
		SetRate(rate).
		SetMintAccount(mint).
		SetAuthorityAccount(authority)
}

// SetRate sets the "rate" parameter.
func (obj *UpdateRate) SetRate(rate int16) *UpdateRate {
	obj.Rate = &rate
	return obj
}

// SetMintAccount sets the "mint" parameter.
// The mint
func (obj *UpdateRate) SetMintAccount(mint common.PublicKey) *UpdateRate {
	obj.AccountMetaSlice[0] = common.Meta(mint).WRITE()
	return obj
}

// GetMintAccount gets the "mint" parameter.
// The mint
func (obj *UpdateRate) GetMintAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(0)
}

// SetAuthorityAccount sets the "authority" parameter.
// The mint rate authority.
func (obj *UpdateRate) SetAuthorityAccount(authority common.PublicKey, multiSigners ...common.PublicKey) *UpdateRate {
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
// The mint rate authority.
func (obj *UpdateRate) GetAuthorityAccount() *common.AccountMeta {
	return obj.AccountMetaSlice.Get(1)
}

func (obj *UpdateRate) SetProgramId(programId *common.PublicKey) *UpdateRate {
	obj._programId = programId
	return obj
}

func (obj *UpdateRate) Build() *Instruction {
	return &Instruction{
		BaseVariant: binary.BaseVariant{
			Impl:   obj,
			TypeID: binary.TypeIDFromBytes([]byte{Instruction_UpdateRate}),
		},
		programId: obj._programId,
		typeIdLen: 1,
	}
}

func (obj *UpdateRate) Validate() error {
	if obj.Rate == nil {
		return errors.New("[UpdateRate] rate param is not set")
	}

	if obj.AccountMetaSlice[0] == nil {
		return errors.New("[UpdateRate] accounts.mint is not set")
	}
	if obj.AccountMetaSlice[1] == nil {
		return errors.New("[UpdateRate] accounts.authority is not set")
	}
	return nil
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (obj *UpdateRate) ValidateAndBuild() (*Instruction, error) {
	if err := obj.Validate(); err != nil {
		return nil, err
	}
	return obj.Build(), nil
}

func (obj *UpdateRate) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.Rate); err != nil {
		return err
	}
	return nil
}

func (obj *UpdateRate) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.Rate); err != nil {
		return err
	}
	return nil
}

func (obj *UpdateRate) EncodeToTree(parent treeout.Branches) {
	parent.Child(format.Program(ProgramName, common.As(ProgramID))).
		ParentFunc(func(programBranch treeout.Branches) {
			programBranch.Child(format.Instruction("UpdateRate")).
				ParentFunc(func(instructionBranch treeout.Branches) {
					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch treeout.Branches) {
						paramsBranch.Child(format.Param("Rate", *obj.Rate))
					})
					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=2]").ParentFunc(func(accountsBranch treeout.Branches) {
						accountsBranch.Child(common.FormatMeta("     mint", obj.AccountMetaSlice.Get(0)))
						accountsBranch.Child(common.FormatMeta("authority", obj.AccountMetaSlice.Get(1)))
					})
				})
		})
}
