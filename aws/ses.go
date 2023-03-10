package aws

import (
	"bytes"
	"context"
	"log"
	"mime/multipart"
	"net/textproto"
	"os"
	"strings"

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
	Cc      string
	Bcc     string
	Subject string
	Data    types.RawMessage
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
func (e *SESEmail) BuildMessage(to, cc, bcc, subject, html, text, fileName string, file []byte) {
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject

	// Create a new multipart writer
	buf := new(bytes.Buffer)
	writer := multipart.NewWriter(buf)

	// Add email headers
	h := make(textproto.MIMEHeader)
	h.Set("From", from)
	h.Set("To", to)
	h.Set("CC", cc)
	h.Set("Return-Path", from)
	h.Set("Subject", subject)
	h.Set("Content-Language", "en-US")
	h.Set("Content-Type", "multipart/mixed; boundary=\""+writer.Boundary()+"\"")
	h.Set("MIME-Version", "1.0")
	_, err := writer.CreatePart(h)
	if err != nil {
		log.Println("Create header part error: ", err.Error())
	}

	// Text body
	h = make(textproto.MIMEHeader)
	h.Set("Content-Transfer-Encoding", "8bit")
	h.Set("Content-Type", "text/plain; charset=utf-8")
	part, err := writer.CreatePart(h)
	if err != nil {
		log.Println("Create text part error: ", err.Error())
	}
	_, err = part.Write([]byte(text))
	if err != nil {
		log.Println("Create text part error: ", err.Error())
	}

	// HTML body
	h = make(textproto.MIMEHeader)
	h.Set("Content-Transfer-Encoding", "quoted-printable")
	h.Set("Content-Type", "text/html; charset=utf-8")
	part, err = writer.CreatePart(h)
	if err != nil {
		log.Println("Create html part error: ", err.Error())
	}
	_, err = part.Write([]byte(html))
	if err != nil {
		log.Println("Create html part error: ", err.Error())
	}

	// File Attachment
	h = make(textproto.MIMEHeader)
	h.Set("Content-Disposition", "attachment; filename="+fileName)
	h.Set("Content-Type", "application/pdf; x-unix-mode=0644; name=\""+fileName+"\"")
	h.Set("Content-Transfer-Encoding", "quoted-printable")
	part, err = writer.CreatePart(h)
	if err != nil {
		log.Println("Create file attachment part error: ", err.Error())
	}
	_, err = part.Write(file)
	if err != nil {
		log.Println("Create file attachment part error: ", err.Error())
	}

	// Close multipart writer
	err = writer.Close()
	if err != nil {
		log.Println("Close multipart writer error: ", err.Error())
	}

	// Get email content
	s := buf.String()
	if strings.Count(s, "\n") < 2 {
		log.Println("Error: invalid e-mail content")
	}
	// s = strings.SplitN(s, "\n", 2)[1]

	e.Data = types.RawMessage{
		Data: []byte(s),
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
	return nil
}
