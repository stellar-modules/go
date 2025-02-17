package resourceadapter

import (
	"context"

	protocol "github.com/stellar-modules/go/sdk/protocols/horizon"
	"github.com/stellar-modules/go/services/horizon/internal/httpx"
	"github.com/stellar-modules/go/services/horizon/internal/txsub"
	"github.com/stellar-modules/go/sdk/support/render/hal"
)

// Populate fills out the details
func PopulateTransactionSuccess(ctx context.Context, dest *protocol.TransactionSuccess, result txsub.Result) {
	dest.Hash = result.Hash
	dest.Ledger = result.LedgerSequence
	dest.Env = result.EnvelopeXDR
	dest.Result = result.ResultXDR
	dest.Meta = result.ResultMetaXDR

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	dest.Links.Transaction = lb.Link("/transactions", result.Hash)
}
