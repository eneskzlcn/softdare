package config

type RabbitMQ struct {
	Username string   `mapstructure:"username"`
	Password string   `mapstructure:"password"`
	Host     string   `mapstructure:"host"`
	Port     string   `mapstructure:"port"`
	Queues   []string `mapstructure:"queues"`
}
