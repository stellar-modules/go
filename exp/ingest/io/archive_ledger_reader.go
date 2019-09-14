package io

import "github.com/stellar/go/sdk/xdr"

// ArchiveLedgerReader placeholder
type ArchiveLedgerReader interface {
	GetSequence() uint32
	Read() (bool, xdr.Transaction, xdr.TransactionResult, error)
}
