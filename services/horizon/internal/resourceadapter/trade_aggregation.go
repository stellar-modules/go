package resourceadapter

import (
	"context"

	"github.com/stellar-modules/go/sdk/amount"
	"github.com/stellar-modules/go/sdk/price"
	protocol "github.com/stellar-modules/go/sdk/protocols/horizon"
	"github.com/stellar-modules/go/services/horizon/internal/db2/history"
)

// Populate fills out the details of a trade using a row from the history_trades
// table.
func PopulateTradeAggregation(
	ctx context.Context,
	dest *protocol.TradeAggregation,
	row history.TradeAggregation,
) error {
	var err error
	dest.Timestamp = row.Timestamp
	dest.TradeCount = row.TradeCount
	dest.BaseVolume, err = amount.IntStringToAmount(row.BaseVolume)
	if err != nil {
		return err
	}
	dest.CounterVolume, err = amount.IntStringToAmount(row.CounterVolume)
	if err != nil {
		return err
	}
	dest.Average = price.StringFromFloat64(row.Average)
	dest.High = row.High.String()
	dest.HighR = row.High
	dest.Low = row.Low.String()
	dest.LowR = row.Low
	dest.Open = row.Open.String()
	dest.OpenR = row.Open
	dest.Close = row.Close.String()
	dest.CloseR = row.Close
	return nil
}
