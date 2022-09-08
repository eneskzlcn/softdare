package config

import "github.com/spf13/viper"

type Config struct {
	Db      DB      `mapstructure:"db"`
	Server  Server  `mapstructure:"server"`
	Session Session `mapstructure:"session"`
}

func LoadConfig(path, name, configType string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(configType)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

var EnvConfigNameMap = map[string]string{
	"":     "local",
	"test": "test",
}
