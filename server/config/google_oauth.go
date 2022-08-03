package config

type GoogleOAuth struct {
	ClientID     string `mapstructure:"clientId" json:"clientId" yaml:"clientId"`
	ClientSecret string `mapstructure:"clientSecret" json:"clientSecret" yaml:"clientSecret"`
	RedirectURL  string `mapstructure:"redirectUrl" json:"redirectUrl" yaml:"redirectUrl"`
}
