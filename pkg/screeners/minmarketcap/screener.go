package minmarketcap

import (
	"context"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

var _ pipeline.PipelineStep = (*MinMarketCap)(nil)

type MinMarketCap struct {
	Configuration Configuration
	Client        *polygon.Client
}

func New(_ context.Context, config Configuration) (*MinMarketCap, error) {
	polygonClient := polygon.New(config.Polygon.Auth.APIKey)

	return &MinMarketCap{
		Configuration: config,
		Client:        polygonClient,
	}, nil
}

func (s MinMarketCap) Description() string {
	return "A simple screener that removes stocks below a configured market cap"
}

func (s MinMarketCap) Process(ctx context.Context, input []types.Equity) ([]types.Equity, error) {
	var returnVal []types.Equity

	for _, val := range input {
		params := models.GetTickerDetailsParams{
			Ticker: val.Ticker,
		}

		details, err := s.Client.GetTickerDetails(ctx, &params)
		if err != nil {
			return nil, err
		}

		if details.Results.MarketCap >= s.Configuration.Minimum {
			val.MarketCap = details.Results.MarketCap
			returnVal = append(returnVal, val)
		}
	}

	return returnVal, nil
}
