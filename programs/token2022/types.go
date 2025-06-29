// Copyright 2021 github.com/gagliardetto
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package token2022

import (
	ag_binary "github.com/gagliardetto/binary"
)

type AuthorityType ag_binary.BorshEnum

const (
	// Authority to mint new tokens
	AuthorityMintTokens AuthorityType = iota

	// Authority to freeze any account associated with the Mint
	AuthorityFreezeAccount

	// Owner of a given token account
	AuthorityAccountOwner

	// Authority to close a token account
	AuthorityCloseAccount
)

type AccountState ag_binary.BorshEnum

const (
	// Account is not yet initialized
	Uninitialized AccountState = iota

	// Account is initialized; the account owner and/or delegate may perform permitted operations
	// on this account
	Initialized

	// Account has been frozen by the mint freeze authority. Neither the account owner nor
	// the delegate are able to perform operations on this account.
	Frozen
)

// ExtensionType Enum
type ExtensionType ag_binary.BorshEnum

const (
	ExtensionTypeUninitialized ExtensionType = iota
	ExtensionTypeTransferFeeConfig
	ExtensionTypeTransferFeeAmount
	ExtensionTypeMintCloseAuthority
	ExtensionTypeConfidentialTransferMint
	ExtensionTypeConfidentialTransferAccount
	ExtensionTypeDefaultAccountState
	ExtensionTypeImmutableOwner
	ExtensionTypeMemoTransfer
	ExtensionTypeNonTransferable
	ExtensionTypeInterestBearingConfig
	ExtensionTypeCpiGuard
	ExtensionTypePermanentDelegate
	ExtensionTypeNonTransferableAccount
	ExtensionTypeTransferHook
	ExtensionTypeTransferHookAccount
	ExtensionTypeConfidentialTransferFeeConfig
	ExtensionTypeConfidentialTransferFeeAmount
	ExtensionTypeMetadataPointer
	ExtensionTypeTokenMetadata
	ExtensionTypeGroupPointer
	ExtensionTypeTokenGroup
	ExtensionTypeGroupMemberPointer
	ExtensionTypeTokenGroupMember
)

func (value ExtensionType) String() string {
	switch value {
	case ExtensionTypeUninitialized:
		return "Uninitialized"
	case ExtensionTypeTransferFeeConfig:
		return "TransferFeeConfig"
	case ExtensionTypeTransferFeeAmount:
		return "TransferFeeAmount"
	case ExtensionTypeMintCloseAuthority:
		return "MintCloseAuthority"
	case ExtensionTypeConfidentialTransferMint:
		return "ConfidentialTransferMint"
	case ExtensionTypeConfidentialTransferAccount:
		return "ConfidentialTransferAccount"
	case ExtensionTypeDefaultAccountState:
		return "DefaultAccountState"
	case ExtensionTypeImmutableOwner:
		return "ImmutableOwner"
	case ExtensionTypeMemoTransfer:
		return "MemoTransfer"
	case ExtensionTypeNonTransferable:
		return "NonTransferable"
	case ExtensionTypeInterestBearingConfig:
		return "InterestBearingConfig"
	case ExtensionTypeCpiGuard:
		return "CpiGuard"
	case ExtensionTypePermanentDelegate:
		return "PermanentDelegate"
	case ExtensionTypeNonTransferableAccount:
		return "NonTransferableAccount"
	case ExtensionTypeTransferHook:
		return "TransferHook"
	case ExtensionTypeTransferHookAccount:
		return "TransferHookAccount"
	case ExtensionTypeConfidentialTransferFeeConfig:
		return "ConfidentialTransferFeeConfig"
	case ExtensionTypeConfidentialTransferFeeAmount:
		return "ConfidentialTransferFeeAmount"
	case ExtensionTypeMetadataPointer:
		return "MetadataPointer"
	case ExtensionTypeTokenMetadata:
		return "TokenMetadata"
	case ExtensionTypeGroupPointer:
		return "GroupPointer"
	case ExtensionTypeTokenGroup:
		return "TokenGroup"
	case ExtensionTypeGroupMemberPointer:
		return "GroupMemberPointer"
	case ExtensionTypeTokenGroupMember:
		return "TokenGroupMember"
	default:
		return ""
	}
}
