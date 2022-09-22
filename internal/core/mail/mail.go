package mail

type Service interface {
	SendTextMail(to, subject, content string) error
}
