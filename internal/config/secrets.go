package config

import "fmt"

type Secrets struct {
	MailService MailService `mapstructure:"mailService"`
}

func (s *Secrets) Print() {
	fmt.Println(*s)
}
