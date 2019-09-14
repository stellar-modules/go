module github.com/stellar-modules/go/exp

go 1.12

require (
	github.com/Masterminds/squirrel v0.0.0-20161115235646-20f192218cf5
	github.com/sirupsen/logrus v1.0.6-0.20180530095059-070c81def33f
	github.com/stellar-modules/go/sdk v0.0.0-20190914042115-7850eaa0ddb8
	github.com/stretchr/testify v1.4.0
	golang.org/x/crypto v0.0.0-20190621222207-cc06ce4a13d4
)

replace github.com/stellar-modules/go/sdk => ../sdk
