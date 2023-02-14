package aws

import (
	"context"
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
	To      string
	Cc      string `default:""`
	Bcc     string `default:""`
	Subject string
	Body    string
}

// NewSESEmail returns an initialized SES email context
func NewSESEmail(to, cc, bcc, subject, body string) *SESEmail {
	// Load assumed user role AWS configuration
	cfg := AssumedRoleConfig()
	// Initialize AWS Simple Email Service Client
	client = sesv2.NewFromConfig(cfg)
	return &SESEmail{To: to, Cc: cc, Bcc: bcc, Subject: subject, Body: body}
}

// Send sends an email using AWS SES
func (e *SESEmail) Send() error {
	// Build email context
	input := &sesv2.SendEmailInput{
		FromEmailAddress: &from,
		Destination: &types.Destination{
			ToAddresses: []string{e.To},
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
		println(err.Error())
		return err
	}
	println(output.MessageId)
	return nil
}
