package minmarketcap

type Configuration struct {
	Minimum float64 `yaml:"minimum"`
	Polygon PolygonConfiguration `yaml:"polygon"`
}

type PolygonConfiguration struct {
	Auth AuthConfiguration `yaml:"auth"`
}

type AuthConfiguration struct {
	APIKey string `yaml:"apiKey"`	
}
