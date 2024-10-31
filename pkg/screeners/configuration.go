package screeners

import (
	"github.com/Sackbuoy/trader/pkg/screeners/minmarketcap"
	"github.com/Sackbuoy/trader/pkg/screeners/stringfilter"
)

type Configuration struct {
	StringFilter *stringfilter.Configuration `yaml:"stringfilter"`
	MinMarketCap *minmarketcap.Configuration `yaml:"minmarketcap"`
}
