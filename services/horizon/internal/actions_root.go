package horizon

import (
	"github.com/stellar-modules/go/sdk/protocols/horizon"
	"github.com/stellar-modules/go/services/horizon/internal/actions"
	"github.com/stellar-modules/go/services/horizon/internal/ledger"
	"github.com/stellar-modules/go/services/horizon/internal/resourceadapter"
	"github.com/stellar-modules/go/sdk/support/render/hal"
)

// Interface verification
var _ actions.JSONer = (*RootAction)(nil)

// RootAction provides a summary of the horizon instance and links to various
// useful endpoints
type RootAction struct {
	Action
}

// JSON renders the json response for RootAction
func (action *RootAction) JSON() error {
	var res horizon.Root
	resourceadapter.PopulateRoot(
		action.R.Context(),
		&res,
		ledger.CurrentState(),
		action.App.horizonVersion,
		action.App.coreVersion,
		action.App.config.NetworkPassphrase,
		action.App.currentProtocolVersion,
		action.App.coreSupportedProtocolVersion,
		action.App.config.FriendbotURL,
	)

	hal.Render(action.W, res)
	return action.Err
}
