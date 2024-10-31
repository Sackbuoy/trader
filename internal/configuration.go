package internal

import (
	"github.com/Sackbuoy/trader/pkg/algs"
	"github.com/Sackbuoy/trader/pkg/brokerages"
	"github.com/Sackbuoy/trader/pkg/inputs"
	"github.com/Sackbuoy/trader/pkg/screeners"
)

type Configuration struct {
	Algorithm algs.Configuration       `yaml:"algorithm"`
	Screener  screeners.Configuration  `yaml:"screener"`
	Input     inputs.Configuration     `yaml:"input"`
	Brokerage brokerages.Configuration `yaml:"brokerage"`
}
