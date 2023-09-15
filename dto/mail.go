package dto

type MailReqBody struct {
	Receiver string `json:"receiver"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}
