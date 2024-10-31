package inputs

import (
	"github.com/Sackbuoy/trader/pkg/inputs/example"
	"github.com/Sackbuoy/trader/pkg/inputs/file"
	"github.com/Sackbuoy/trader/pkg/inputs/nyse"
)

type Configuration struct {
	Example *example.Configuration `yaml:"example"`
	NYSE    *nyse.Configuration    `yaml:"nyse"`
	File    *file.Configuration    `yaml:"file"`

	Exclusions []string `yaml:"exclusions"`
	Inclusions []string `yaml:"inclusions"`
}
