module github.com/stellar-modules/go/services/horizon

go 1.12

require (
	bitbucket.org/ww/goautoneg v0.0.0-20120707110453-75cd24fc2f2c
	github.com/Masterminds/squirrel v0.0.0-20161115235646-20f192218cf5
	github.com/elazarl/go-bindata-assetfs v1.0.0
	github.com/getsentry/raven-go v0.0.0-20160805001729-c9d3cc542ad1
	github.com/go-chi/chi v3.1.5+incompatible
	github.com/go-errors/errors v0.0.0-20150906023321-a41850380601
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/guregu/null v2.1.3-0.20151024101046-79c5bd36b615+incompatible
	github.com/hashicorp/golang-lru v0.5.0 // indirect
	github.com/jmoiron/sqlx v1.2.0
	github.com/lib/pq v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/rcrowley/go-metrics v0.0.0-20160113235030-51425a2415d2
	github.com/rs/cors v0.0.0-20160617231935-a62a804a8a00
	github.com/rubenv/sql-migrate v0.0.0-20190717103323-87ce952f7079
	github.com/sebest/xff v0.0.0-20150611211316-7a36e3a787b5
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/sirupsen/logrus v1.0.6-0.20180530095059-070c81def33f
	github.com/spf13/cobra v0.0.0-20160830174925-9c28e4bbd74e
	github.com/spf13/viper v0.0.0-20150621231900-db7ff930a189
	github.com/stellar-modules/go/exp v0.0.0-20190914042115-7850eaa0ddb8
	github.com/stellar-modules/go/sdk v0.0.0-20190914042115-7850eaa0ddb8
	github.com/stellar/throttled v2.2.3-0.20190823235211-89d75816f59d+incompatible
	github.com/stretchr/testify v1.4.0
	golang.org/x/tools v0.0.0-20190624180213-70d37148ca0c // indirect
	gopkg.in/tylerb/graceful.v1 v1.2.13
)

replace github.com/stellar-modules/go/sdk => ../../sdk

replace github.com/stellar-modules/go/exp => ../../exp
