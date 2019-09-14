module github.com/stellar-modules/go/tools/stellar-archivist

go 1.12

require (
	github.com/spf13/cobra v0.0.0-20160830174925-9c28e4bbd74e
	github.com/stellar-modules/go/sdk v0.0.0-20190914042115-7850eaa0ddb8
	github.com/stretchr/testify v1.4.0
)

replace github.com/stellar-modules/go/sdk => ../../sdk
