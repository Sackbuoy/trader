package nyse

import (
	"context"
	"os"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

var _ pipeline.PipelineStep = (*NYSEInput)(nil)

type NYSEInput struct {
	Configuration Configuration
	Client        *polygon.Client
	Inclusions    []string
	Exclusions    []string
}

func New(_ context.Context,
	config Configuration,
	inclusions []string,
	exclusions []string,
) (*NYSEInput, error) {
	polygonClient := polygon.New(os.Getenv("POLYGON_API_KEY"))

	return &NYSEInput{
		Configuration: config,
		Client:        polygonClient,
		Inclusions:    inclusions,
		Exclusions:    exclusions,
	}, nil
}

func (i NYSEInput) Description() string {
	return "Returns all equities currently on the NYSE"
}

func (i NYSEInput) Process(ctx context.Context, input []types.Equity) ([]types.Equity, error) {
	params := models.ListTickersParams{}

	tickers := i.Client.ListTickers(ctx, &params)
	cur := tickers.Item()

	for tickers.Next() {
		if !i.isIncludedOrExcluded(cur.Ticker) {
			input = append(input, types.Equity{
				Ticker: cur.Ticker,
			})
		}

		cur = tickers.Item()
	}

	// make sure inclusions get added
	for _, v := range i.Inclusions {
		input = append(input, types.Equity{
			Ticker: v,
		})
	}

	return input, nil
}

func (i NYSEInput) isIncludedOrExcluded(ticker string) bool {
	switch {
	case contains(ticker, i.Inclusions):
		return true
	case contains(ticker, i.Exclusions):
		return true
	}
	return false
}

func contains(val string, arr []string) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
