package resourceadapter

import (
	"context"

	protocol "github.com/stellar-modules/go/sdk/protocols/horizon"
	"github.com/stellar-modules/go/sdk/xdr"
)

func PopulateAsset(ctx context.Context, dest *protocol.Asset, asset xdr.Asset) error {
	return asset.Extract(&dest.Type, &dest.Code, &dest.Issuer)
}
