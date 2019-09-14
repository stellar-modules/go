module github.com/stellar-modules/go/services/federation

go 1.12

require (
	github.com/go-chi/chi v3.1.5+incompatible
	github.com/spf13/cobra v0.0.0-20160830174925-9c28e4bbd74e
	github.com/stellar-modules/go/sdk v0.0.0-20190914042509-df015cd57c57
)

replace github.com/stellar-modules/go/sdk => ../../sdk
