package config

type App struct {
	Address    string `mapstructure:"address"`
	SessionKey string `mapstructure:"sessionKey"`
}
