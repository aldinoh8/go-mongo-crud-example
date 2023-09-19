package services

type Mailer interface {
	SendMail(string, string, string) error
}
