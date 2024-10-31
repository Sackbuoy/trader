package brokerages

import "github.com/Sackbuoy/trader/pkg/brokerages/tradier"

type Configuration struct {
	Tradier *tradier.Configuration `yaml:"tradier"`
}
