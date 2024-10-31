package stringfilter

import (
	"context"
	"strings"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
)

var _ pipeline.PipelineStep = (*StringFilter)(nil)

type StringFilter struct {
	Configuration Configuration
}

func New(_ context.Context,
	config Configuration,
) (*StringFilter, error) {
	return &StringFilter{
		Configuration: config,
	}, nil
}

func (s StringFilter) Description() string {
	return "A simple screener that removes stocks with Tickers that match a given filter"
}

func (s StringFilter) Process(_ context.Context, input []types.Equity) ([]types.Equity, error) {
	var returnVal []types.Equity

	for _, v := range input {
		if !strings.Contains(v.Ticker, s.Configuration.Filter) {
			returnVal = append(returnVal, v)
		}
	}

	return returnVal, nil
}
