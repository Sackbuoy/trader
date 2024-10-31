package tradier

import (
	"context"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
)

var _ pipeline.PipelineStep = (*TradierBroker)(nil)

type TradierBroker struct {
	Configuration Configuration
	Client        *Client
}

func New(ctx context.Context, config Configuration) (*TradierBroker, error) {
	client, err := CreateTradierClient(ctx, config)
	if err != nil {
		return nil, err
	}

	return &TradierBroker{
		Configuration: config,
		Client:        client,
	}, nil
}

func (i TradierBroker) Description() string {
	return "Brokerage Integration for Tradier"
}

func (i TradierBroker) Process(ctx context.Context, input []types.Equity) ([]types.Equity, error) {
	for _, v := range input {

		switch v.Action {
		case "BUY":
			err := i.Client.PlaceEquityOrder(ctx, v)
			if err != nil {
				return nil, err
			}
		default:
			continue
		}
	}

	return input, nil
}
