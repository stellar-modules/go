package resourceadapter

import (
	protocol "github.com/stellar-modules/go/sdk/protocols/horizon"
	"github.com/stellar-modules/go/services/horizon/internal/db2/core"
)

func PopulateAccountFlags(dest *protocol.AccountFlags, row core.Account) {
	dest.AuthRequired = row.IsAuthRequired()
	dest.AuthRevocable = row.IsAuthRevocable()
	dest.AuthImmutable = row.IsAuthImmutable()
}
