package handlers

import (
	"github.com/stellar-modules/go/sdk/clients/federation"
	hc "github.com/stellar-modules/go/sdk/clients/horizonclient"
	"github.com/stellar-modules/go/sdk/clients/stellartoml"
	"github.com/stellar-modules/go/services/bridge/internal/config"
	"github.com/stellar-modules/go/services/bridge/internal/db"
	"github.com/stellar-modules/go/services/bridge/internal/listener"
	"github.com/stellar-modules/go/services/bridge/internal/submitter"
	"github.com/stellar-modules/go/sdk/support/http"
)

// RequestHandler implements bridge server request handlers
type RequestHandler struct {
	Config               *config.Config                          `inject:""`
	Client               http.SimpleHTTPClientInterface          `inject:""`
	Horizon              hc.ClientInterface                      `inject:""`
	Database             db.Database                             `inject:""`
	StellarTomlResolver  stellartoml.ClientInterface             `inject:""`
	FederationResolver   federation.ClientInterface              `inject:""`
	TransactionSubmitter submitter.TransactionSubmitterInterface `inject:""`
	PaymentListener      *listener.PaymentListener               `inject:""`
}
