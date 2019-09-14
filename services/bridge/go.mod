module github.com/stellar-modules/go/services/bridge

go 1.12

require (
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/inject v0.0.0-20161006174721-cc1aa653e50f
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/structtag v0.0.0-20150214074306-217e25fb9691 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/go-chi/chi v3.1.5+incompatible
	github.com/rubenv/sql-migrate v0.0.0-20190717103323-87ce952f7079
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/sirupsen/logrus v1.0.6-0.20180530095059-070c81def33f
	github.com/spf13/cobra v0.0.0-20160830174925-9c28e4bbd74e
	github.com/stellar-modules/go/sdk v0.0.0-20190914042115-7850eaa0ddb8
	github.com/stellar-modules/go/services/internal v0.0.0-20190914042115-7850eaa0ddb8
	github.com/stretchr/testify v1.4.0
	golang.org/x/tools v0.0.0-20190624180213-70d37148ca0c // indirect
)

replace github.com/stellar-modules/go/sdk => ../../sdk

replace github.com/stellar-modules/go/services/internal => ../internal
