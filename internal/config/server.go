package config

type Server struct {
	Address    string `mapstructure:"address"`
	SessionKey string `mapstructure:"sessionKey"`
}
