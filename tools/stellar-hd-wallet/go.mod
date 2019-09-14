module github.com/stellar/go/tools/stellar-hd-wallet

go 1.12

require (
	github.com/spf13/cobra v0.0.0-20160830174925-9c28e4bbd74e
	github.com/stellar/go/exp v0.0.0-00010101000000-000000000000
	github.com/stellar/go/sdk v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.4.0
	github.com/tyler-smith/go-bip39 v0.0.0-20180618194314-52158e4697b8
)

replace github.com/stellar/go/sdk => ../../sdk

replace github.com/stellar/go/exp => ../../exp
