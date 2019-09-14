module github.com/stellar/go/services/keystore

go 1.12

require (
	github.com/lib/pq v1.2.0
	github.com/rubenv/sql-migrate v0.0.0-20190717103323-87ce952f7079
	github.com/sirupsen/logrus v1.0.6-0.20180530095059-070c81def33f
	github.com/stellar/go/sdk v0.0.0-00010101000000-000000000000
)

replace github.com/stellar/go/sdk => ../../sdk
