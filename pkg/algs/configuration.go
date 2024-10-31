package algs

import avgcompare "github.com/Sackbuoy/trader/pkg/algs/avgCompare"

type Configuration struct {
	AvgCompare *avgcompare.Configuration `yaml:"avgCompare"`
}
