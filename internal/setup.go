package internal

import (
	"context"

	"github.com/Sackbuoy/trader/internal/pipeline"
	"github.com/Sackbuoy/trader/pkg/algs"
	avgcompare "github.com/Sackbuoy/trader/pkg/algs/avgCompare"
	"github.com/Sackbuoy/trader/pkg/brokerages"
	"github.com/Sackbuoy/trader/pkg/brokerages/tradier"
	"github.com/Sackbuoy/trader/pkg/inputs"
	"github.com/Sackbuoy/trader/pkg/inputs/example"
	"github.com/Sackbuoy/trader/pkg/inputs/file"
	"github.com/Sackbuoy/trader/pkg/inputs/nyse"
	"github.com/Sackbuoy/trader/pkg/screeners"
	"github.com/Sackbuoy/trader/pkg/screeners/minmarketcap"
	"github.com/Sackbuoy/trader/pkg/screeners/stringfilter"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func BuildConfiguration(
	viperRef *viper.Viper,
	config *Configuration,
	envVars ...string,
) error {
	for _, envVar := range envVars {
		_ = viperRef.BindEnv(envVar)
	}

	initialSetup(viperRef)

	err := viperRef.ReadInConfig()
	if err != nil {
		return errors.Wrap(err, "failed to read in configuration")
	}

	err = viperRef.Unmarshal(&config)
	if err != nil {
		return errors.Wrap(err, "failed to unmarshal configuration into expected structure")
	}

	return nil
}

func initialSetup(viperRef *viper.Viper) {
	viperRef.SetConfigName("configuration")
	viperRef.SetConfigType("yaml")
	viperRef.AddConfigPath(".")

	viperRef.AutomaticEnv()
}

func SetupPipeline(ctx context.Context, config Configuration) (*pipeline.Pipeline, error) {
	input, err := setupInput(ctx, config.Input)
	if err != nil {
		return nil, err
	}

	screener, err := setupScreener(ctx, config.Screener)
	if err != nil {
		return nil, err
	}

	alg, err := setupAlgorithm(ctx, config.Algorithm)
	if err != nil {
		return nil, err
	}

	brokerage, err := setupBrokerage(ctx, config.Brokerage)
	if err != nil {
		return nil, err
	}

	return &pipeline.Pipeline{
		Input:     input,
		Screener:  screener,
		Algorithm: alg,
		Brokerage: brokerage,
	}, nil
}

func setupInput(ctx context.Context, config inputs.Configuration) (pipeline.PipelineStep, error) {
	var (
		step pipeline.PipelineStep
		err  error
	)

	switch {
	case config.Example != nil:
		step, err = example.New(ctx, *config.Example, config.Inclusions, config.Exclusions)
	case config.NYSE != nil:
		step, err = nyse.New(ctx, *config.NYSE, config.Inclusions, config.Exclusions)
	case config.File != nil:
		step, err = file.New(ctx, *config.File, config.Inclusions, config.Exclusions)
	default:
		err = errors.New("Invalid configuration for Input")
	}

	return step, err
}

func setupScreener(ctx context.Context, config screeners.Configuration) (pipeline.PipelineStep, error) {
	var (
		step pipeline.PipelineStep
		err  error
	)

	switch {
	case config.StringFilter != nil:
		step, err = stringfilter.New(ctx, *config.StringFilter)
	case config.MinMarketCap != nil:
		step, err = minmarketcap.New(ctx, *config.MinMarketCap)
	default:
		err = errors.New("Invalid configuration for Screener")
	}

	return step, err
}

func setupAlgorithm(ctx context.Context, config algs.Configuration) (pipeline.PipelineStep, error) {
	var (
		step pipeline.PipelineStep
		err  error
	)

	switch {
	case config.AvgCompare != nil:
		step, err = avgcompare.New(ctx, *config.AvgCompare)
	default:
		err = errors.New("Invalid configuration for Algorithm")
	}

	return step, err
}

func setupBrokerage(ctx context.Context, config brokerages.Configuration) (pipeline.PipelineStep, error) {
	var (
		step pipeline.PipelineStep
		err  error
	)

	switch {
	case config.Tradier != nil:
		step, err = tradier.New(ctx, *config.Tradier)
	default:
		err = errors.New("Invalid configuration for Brokerage")
	}

	return step, err
}
