package pipeline

import (
	"context"

	"github.com/Sackbuoy/trader/internal/types"
)

type Pipeline struct {
	Input     PipelineStep
	Screener  PipelineStep
	Algorithm PipelineStep
	Brokerage PipelineStep
}

func New(input, screener, algorithm, brokerage PipelineStep) (*Pipeline, error) {
	return &Pipeline{
		Input:     input,
		Screener:  screener,
		Algorithm: algorithm,
		Brokerage: brokerage,
	}, nil
}

func (p *Pipeline) Run(ctx context.Context) error {
	var (
		equityList []types.Equity
		err        error
	)

	equityList, err = p.Input.Process(ctx, equityList)
	if err != nil {
		return err
	}

	equityList, err = p.Screener.Process(ctx, equityList)
	if err != nil {
		return err
	}

	equityList, err = p.Algorithm.Process(ctx, equityList)
	if err != nil {
		return err
	}

	equityList, err = p.Brokerage.Process(ctx, equityList)
	if err != nil {
		return err
	}

	displayResults(equityList)

	return nil
}

func displayResults(values []types.Equity) {
	for _, v := range values {
		v.Print()
	}
}
