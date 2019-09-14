module github.com/stellar/go/tools/horizon-verify

go 1.12

require (
	github.com/spf13/cobra v0.0.0-20160830174925-9c28e4bbd74e
	github.com/stellar/go/sdk v0.0.0-00010101000000-000000000000
)

replace github.com/stellar/go/sdk => ../../sdk
