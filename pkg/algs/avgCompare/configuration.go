package avgcompare

import "time"

type Configuration struct {
	Short time.Duration `yaml:"short"`
	Long  time.Duration `yaml:"long"`
}
