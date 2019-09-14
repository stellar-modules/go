package ticker

import (
	"github.com/stellar-modules/go/services/ticker/internal/gql"
	"github.com/stellar-modules/go/services/ticker/internal/tickerdb"
	hlog "github.com/stellar-modules/go/sdk/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
