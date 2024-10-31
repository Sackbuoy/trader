package file

import (
	"bufio"
	"context"
	"os"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/internal/types"
	polygon "github.com/polygon-io/client-go/rest"
)

var _ pipeline.PipelineStep = (*FileInput)(nil)

type FileInput struct {
	Configuration Configuration
	Client        *polygon.Client
	File          *os.File
	Inclusions    []string
	Exclusions    []string
}

func New(_ context.Context,
	config Configuration,
	inclusions []string,
	exclusions []string,
) (*FileInput, error) {
	return &FileInput{
		Configuration: config,
		Inclusions:    inclusions,
		Exclusions:    exclusions,
	}, nil
}

func (i FileInput) Description() string {
	return "Adds tickers from a given file"
}

func (i FileInput) Process(_ context.Context, input []types.Equity) ([]types.Equity, error) {
	file, err := os.Open(i.Configuration.Path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ticker := scanner.Text()
		if !i.isIncludedOrExcluded(ticker) {
			input = append(input, types.Equity{
				Ticker: ticker,
			})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for _, val := range i.Inclusions {
		input = append(input, types.Equity{
			Ticker: val,
		})
	}

	return input, nil
}

func (i FileInput) isIncludedOrExcluded(ticker string) bool {
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
