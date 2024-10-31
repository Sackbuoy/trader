package example

import (
	"context"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
)

var _ pipeline.PipelineStep = (*ExampleInput)(nil)

type ExampleInput struct {
	Configuration Configuration
	Inclusions    []string
	Exclusions    []string
}

func New(_ context.Context,
	config Configuration,
	inclusions []string,
	exclusions []string,
) (*ExampleInput, error) {
	return &ExampleInput{
		Configuration: config,
		Inclusions:    inclusions,
		Exclusions:    exclusions,
	}, nil
}

func (i ExampleInput) Description() string {
	return "A simple Example that returns a short list of equities"
}

func (i ExampleInput) Process(_ context.Context, input []types.Equity) ([]types.Equity, error) {
	newVals := []string{
		"AAPL",
		"MSFT",
		"GOOG",
		"META",
	}

	for _, v := range newVals {
		input = append(input, types.Equity{
			Ticker: v,
		})
	}
	return input, nil
}
