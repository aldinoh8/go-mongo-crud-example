package services

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type Mailer struct {
	Url    string
	Client http.Client
}

func NewMailer(serviceUrl string) Mailer {
	return Mailer{
		Url:    serviceUrl,
		Client: *http.DefaultClient,
	}
}

func (m Mailer) SendMail(receiver, subject, body string) error {
	url := fmt.Sprintf("%s/mail", m.Url)
	reqBody := fmt.Sprintf(`{
    "receiver": "%s",
    "subject": "%s",
    "body": "%s"
	}`, receiver, subject, body)

	req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))

	response, err := m.Client.Do(req)
	if err != nil {
		return errors.New("failed to hit Shipping API")
	}
	defer response.Body.Close()
	return nil
}
