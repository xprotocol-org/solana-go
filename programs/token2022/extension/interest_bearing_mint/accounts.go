// This code was AUTOGENERATED using the library.
// Please DO NOT EDIT THIS FILE.

package interest_bearing_mint

import (
	common "github.com/gagliardetto/solana-go/programs/common"
	binary "github.com/gagliardetto/binary"
	solanago "github.com/gagliardetto/solana-go"
)

// InterestBearingConfig Struct
type InterestBearingConfig struct {
	// Authority that can set the interest rate and authority
	RateAuthority common.PublicKey
	// Timestamp of initialization, from which to base interest calculations
	InitializationTimestamp solanago.UnixTimeSeconds
	// Average rate from initialization until the last time it was updated
	PreUpdateAverageRate int16
	// Timestamp of the last update, used to calculate the total amount accrued
	LastUpdateTimestamp solanago.UnixTimeSeconds
	// Current rate, since the last update
	CurrentRate int16
}

func (obj *InterestBearingConfig) MarshalWithEncoder(encoder *binary.Encoder) (err error) {
	if err = encoder.Encode(&obj.RateAuthority); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.InitializationTimestamp); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.PreUpdateAverageRate); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.LastUpdateTimestamp); err != nil {
		return err
	}
	if err = encoder.Encode(&obj.CurrentRate); err != nil {
		return err
	}
	return nil
}

func (obj *InterestBearingConfig) UnmarshalWithDecoder(decoder *binary.Decoder) (err error) {
	if err = decoder.Decode(&obj.RateAuthority); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.InitializationTimestamp); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.PreUpdateAverageRate); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.LastUpdateTimestamp); err != nil {
		return err
	}
	if err = decoder.Decode(&obj.CurrentRate); err != nil {
		return err
	}
	return nil
}
