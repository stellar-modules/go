package horizon

import (
	"net/http"

	hProblem "github.com/stellar-modules/go/services/horizon/internal/render/problem"
	"github.com/stellar-modules/go/sdk/support/render/problem"
)

// RateLimitExceededAction renders a 429 response
type RateLimitExceededAction struct {
	Action
}

func (action RateLimitExceededAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ap := &action.Action
	ap.Prepare(w, r)
	problem.Render(action.R.Context(), action.W, hProblem.RateLimitExceeded)
}
