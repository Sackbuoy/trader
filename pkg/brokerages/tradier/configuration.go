package tradier

// import "net/url"

type Configuration struct {
	URL string `yaml:"url"`
	Auth AuthConfig `yaml:"auth"`
	AccountID string `yaml:"accountID"`
}

type AuthConfig struct {
	AccessToken string `yaml:"accessToken"`
}
