package services

type MailerInterface interface {
	SendMail(string, string, string) error
}
