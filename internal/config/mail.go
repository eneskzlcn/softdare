package config

type MailService struct {
	SenderMail     string `mapstructure:"senderMail"`
	SenderPassword string `mapstructure:"senderPassword"`
	Host           string `mapstructure:"host"`
	Port           int    `mapstructure:"port"`
}
