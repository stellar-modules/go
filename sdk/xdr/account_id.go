package xdr

import (
	"errors"
	"fmt"

	"github.com/stellar-modules/go/sdk/strkey"
)

// Address returns the strkey encoded form of this AccountId.  This method will
// panic if the accountid is backed by a public key of an unknown type.
func (aid *AccountId) Address() string {
	if aid == nil {
		return ""
	}

	switch aid.Type {
	case PublicKeyTypePublicKeyTypeEd25519:
		ed := aid.MustEd25519()
		raw := make([]byte, 32)
		copy(raw, ed[:])
		return strkey.MustEncode(strkey.VersionByteAccountID, raw)
	default:
		panic(fmt.Errorf("Unknown account id type: %v", aid.Type))
	}
}

// Equals returns true if `other` is equivalent to `aid`
func (aid *AccountId) Equals(other AccountId) bool {
	if aid.Type != other.Type {
		return false
	}

	switch aid.Type {
	case PublicKeyTypePublicKeyTypeEd25519:
		l := aid.MustEd25519()
		r := other.MustEd25519()
		return l == r
	default:
		panic(fmt.Errorf("Unknown account id type: %v", aid.Type))
	}
}

// LedgerKey implements the `Keyer` interface
func (aid *AccountId) LedgerKey() (ret LedgerKey) {
	err := ret.SetAccount(*aid)
	if err != nil {
		panic(err)
	}

	return
}

// MarshalBinaryCompress marshals AccountId to []byte but unlike
// MarshalBinary() it removes all unnecessary bytes, exploting the fact
// that XDR is padding data to 4 bytes in union discriminants etc.
// It's primary use is in ingest/io.StateReader that keep LedgerKeys in
// memory so this function decrease memory requirements.
//
// Warning, do not use UnmarshalBinary() on data encoded using this method!
func (aid AccountId) MarshalBinaryCompress() ([]byte, error) {
	m := []byte{byte(aid.Type)}

	switch aid.Type {
	case PublicKeyTypePublicKeyTypeEd25519:
		pk, err := aid.Ed25519.MarshalBinary()
		if err != nil {
			return nil, err
		}
		m = append(m, pk...)
	default:
		panic("Unknown type")
	}

	return m, nil
}

func MustAddress(address string) AccountId {
	aid := AccountId{}
	err := aid.SetAddress(address)
	if err != nil {
		panic(err)
	}
	return aid
}

// SetAddress modifies the receiver, setting it's value to the AccountId form
// of the provided address.
func (aid *AccountId) SetAddress(address string) error {
	if aid == nil {
		return nil
	}

	raw, err := strkey.Decode(strkey.VersionByteAccountID, address)
	if err != nil {
		return err
	}

	if len(raw) != 32 {
		return errors.New("invalid address")
	}

	var ui Uint256
	copy(ui[:], raw)

	*aid, err = NewAccountId(PublicKeyTypePublicKeyTypeEd25519, ui)

	return err
}
