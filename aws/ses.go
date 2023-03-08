package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
)

var (
	client *sesv2.Client
	from   = os.Getenv("SMTP_FROM")
)

// SESEmail represents an email context
type SESEmail struct {
	To      string `json:"to"`
	Cc      string `json:"cc" default:""`
	Bcc     string `json:"bcc" default:""`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

// NewSESEmail returns an initialized SES email context.
func NewSESEmail() *SESEmail {
	// Load assumed user role AWS configuration
	cfg := AssumedRoleConfig()
	// Initialize AWS Simple Email Service Client
	client = sesv2.NewFromConfig(cfg)
	return &SESEmail{}
}

// BuildMessage builds email context.
func (e *SESEmail) BuildMessage(to, cc, bcc, subject, body string) {
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.Body = body
}

// Send sends an email using AWS SES.
func (e *SESEmail) Send() error {
	// Input email content
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &from,
		Destination: &types.Destination{
			ToAddresses:  []string{e.To},
			BccAddresses: []string{e.Bcc},
		},
		Content: &types.EmailContent{
			Simple: &types.Message{
				Subject: &types.Content{
					Data: &e.Subject,
				},
				Body: &types.Body{
					Text: &types.Content{
						Data: &e.Body,
					},
				},
			},
		},
	}

	// Send email
	output, err := client.SendEmail(context.Background(), input)
	if err != nil {
		log.Println("Send email error: ", err.Error())
		return err
	}
	log.Println("Email sent: ", *output.MessageId)
	return nil
}
