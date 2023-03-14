package aws

import (
	"bytes"
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"gopkg.in/gomail.v2"
)

var (
	client *sesv2.Client
	from   = os.Getenv("SMTP_FROM")
)

// SESEmail represents an email context
type SESEmail struct {
	To       string
	Cc       string
	Bcc      string
	Data     types.RawMessage
	FileName string
}

// SESEmailClient initializes the AWS SES client.
func SESEmailClient() {
	// Load assumed user role AWS configuration
	cfg := AssumedRoleConfig()
	// Initialize AWS Simple Email Service Client
	client = sesv2.NewFromConfig(cfg)
}

// NewSESEmail returns an initialized SES email context.
func NewSESEmail() *SESEmail {
	return &SESEmail{}
}

// BuildMessage builds email context.
func (e *SESEmail) BuildMessage(to, cc, bcc, subject, html, text, fileName string) {
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.FileName = fileName

	// Initialize new email message
	msg := gomail.NewMessage()

	// Set email headers
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)

	// Set email body
	msg.SetBody("text/plain", text)
	msg.SetBody("text/html", html)

	// Set email attachment
	msg.Attach(fileName)

	// Create a new buffer to add raw email data
	var emailRaw bytes.Buffer
	msg.WriteTo(&emailRaw)

	// Set email data
	e.Data = types.RawMessage{
		Data: emailRaw.Bytes(),
	}
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
			Raw: &e.Data,
		},
	}

	// Send email
	output, err := client.SendEmail(context.Background(), input)
	if err != nil {
		log.Println("Send email error: ", err.Error())
		return err
	}
	log.Println("Email sent: ", *output.MessageId)

	// Delete attachment
	os.Remove(e.FileName)
	return nil
}
