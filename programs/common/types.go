package common

import (
	ag_solana "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/text/format"
)

type AccountMeta = ag_solana.AccountMeta
type AccountMetaSlice []*ag_solana.AccountMeta

type PublicKey = ag_solana.PublicKey

func MustPublicKeyFromBase58(input string) PublicKey {
	return ag_solana.MustPublicKeyFromBase58(input)
}

func IsZero(pubkey PublicKey) bool {
	for _, item := range pubkey.Bytes() {
		if item != 0 {
			return false
		}
	}
	return true
}

func As(pubkey PublicKey) ag_solana.PublicKey {
	return ag_solana.PublicKeyFromBytes(pubkey.Bytes())
}

func Meta(
	pubKey PublicKey,
) *ag_solana.AccountMeta {
	return &ag_solana.AccountMeta{
		PublicKey: pubKey,
	}
}

func NewAccountMeta(
	pubKey PublicKey,
	WRITE bool,
	SIGNER bool,
) *AccountMeta {
	return &AccountMeta{
		PublicKey:  pubKey,
		IsWritable: WRITE,
		IsSigner:   SIGNER,
	}
}

func (slice AccountMetaSlice) Get(index int) *AccountMeta {
	if index >= len(slice) {
		return nil
	}
	return slice[index]
}

func (slice *AccountMetaSlice) Append(account *AccountMeta) {
	*slice = append(*slice, account)
}

func (slice *AccountMetaSlice) SetAccounts(accounts []*AccountMeta) error {
	*slice = accounts
	return nil
}

func (slice AccountMetaSlice) GetAccounts() []*AccountMeta {
	out := make([]*AccountMeta, 0, len(slice))
	for i := range slice {
		if slice[i] != nil {
			out = append(out, slice[i])
		}
	}
	return out
}

func FormatMeta(name string, meta *AccountMeta) string {
	var out *ag_solana.AccountMeta
	if meta != nil {
		out = &ag_solana.AccountMeta{
			PublicKey:  ag_solana.PublicKey(meta.PublicKey),
			IsWritable: meta.IsWritable,
			IsSigner:   meta.IsSigner,
		}
	}
	return format.Meta(name, out)
}

func ConvertMeta(input []*ag_solana.AccountMeta) []*AccountMeta {
	var ret []*AccountMeta
	for _, item := range input {
		ret = append(ret, &AccountMeta{
			PublicKey:  PublicKey(item.PublicKey),
			IsSigner:   item.IsSigner,
			IsWritable: item.IsWritable,
		})
	}
	return ret
}

type AccountsSettable interface {
	SetAccounts(accounts []*AccountMeta) error
}

type AccountsGettable interface {
	GetAccounts() (accounts []*AccountMeta)
}
