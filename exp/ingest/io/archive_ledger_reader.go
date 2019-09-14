package io

import "github.com/stellar-modules/go/sdk/xdr"

// ArchiveLedgerReader placeholder
type ArchiveLedgerReader interface {
	GetSequence() uint32
	Read() (bool, xdr.Transaction, xdr.TransactionResult, error)
}
