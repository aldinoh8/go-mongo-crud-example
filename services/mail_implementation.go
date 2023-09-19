package services

import (
	"fmt"
	"net/http"
	"strings"
)

type MailerImplementation struct {
	Url    string
	Client http.Client
}

var mailer *MailerImplementation

func NewMailer(serviceUrl string) *MailerImplementation {
	if mailer == nil {
		mailer = &MailerImplementation{
			Url:    serviceUrl,
			Client: *http.DefaultClient,
		}
	}

	return mailer
}

func (m MailerImplementation) SendMail(receiver, subject, body string) error {
	url := fmt.Sprintf("%s/mail", m.Url)
	reqBody := fmt.Sprintf(`{
		"receiver": "%s",
		"subject": "%s",
		"body": "%s"
	}`, receiver, subject, body)

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
	if err != nil {
		return err
	}

	_, err = m.Client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
