package main

import (
	"context"

	"github.com/Sackbuoy/trader/internal"
	"github.com/spf13/viper"
)

func main() {
	var config internal.Configuration
	ctx := context.Background()

	err := internal.BuildConfiguration(
		viper.New(),
		&config,
	)
	if err != nil {
		panic(err)
	}

	pipeline, err := internal.SetupPipeline(ctx, config)
	if err != nil {
		panic(err)
	}

	err = pipeline.Run(ctx)
	if err != nil {
		panic(err)
	}
}
