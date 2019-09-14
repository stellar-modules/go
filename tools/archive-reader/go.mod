module github.com/stellar/go/tools/archive-reader

go 1.12

require (
	github.com/stellar/go/exp v0.0.0-00010101000000-000000000000
	github.com/stellar/go/sdk v0.0.0-00010101000000-000000000000
)

replace github.com/stellar/go/sdk => ../../sdk

replace github.com/stellar/go/exp => ../../exp
