package email

import (
	"fmt"
	"log"
	"time"

	"github.com/resend/resend-go/v3"
)

type ResendClient struct {
	cli *resend.Client
}

func NewClient(key string) *ResendClient {
	return &ResendClient{
		cli: resend.NewClient(key),
	}

}

func (r *ResendClient) SendEmail(message string) {
	allow(time.Now())
	params := &resend.SendEmailRequest{
		From:    "TYSONCLOUD <notifications@tysonjenkins.dev>",
		To:      []string{"tyson.j.jenkins@gmail.com"},
		Html:    fmt.Sprintf("<h1>%s</h1>", message),
		Subject: "WARNING",
	}

	_, err := r.cli.Emails.Send(params)
	if err != nil {
		log.Println(err)
		return
	}
}

func allow(lastEmail time.Time) bool {
	if lastEmail.Sub(time.Now()) >= time.Hour*24 {
		return true
	}
	return false
}
