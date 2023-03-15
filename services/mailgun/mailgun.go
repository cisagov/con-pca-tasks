package mailgun

import (
	"bytes"
	"context"
	"strconv"
	"text/template"
	"time"

	"github.com/cisagov/con-pca-tasks/database/collections"

	"github.com/mailgun/mailgun-go/v3"
)

// MailgunEmail represents an email context
type MailgunEmail struct {
	From    string
	To      string
	Cc      string
	Bcc     string
	Subject string
	Html    string
	Text    string
	Data    []byte
}

type EmailTags struct {
	URL                  string
	TargetFirstName      string
	TargetPosition       string
	CustomerCity         string
	TimeCurrentDateShort string
	FakeFirstNameMale    string
}

// NewMailgunEmail returns an initialized Mailgun email context.
func NewMailgunEmail() *MailgunEmail {
	return &MailgunEmail{}
}

// BuildMessage builds email context.
func (e *MailgunEmail) BuildMessage(from, to, cc, bcc, subject, text, html string) {
	e.From = from
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.Text = text
	e.Html = html
}

func GetMailgunClient(domain, apiKey string) *mailgun.MailgunImpl {
	client := mailgun.NewMailgun(domain, apiKey)
	return client
}

func FormatTags(tagged_text string, email_tags EmailTags) (string, error) {
	html, err := template.New("email").Parse(tagged_text)
	if err != nil {
		return "error", err
	}
	var pnt bytes.Buffer
	if err := html.Execute(&pnt, email_tags); err != nil {
		return "error", err
	}
	formatted_string := pnt.String()
	return formatted_string, err
}

func GenerateEmailTags(target collections.Target, customer collections.Customer) (EmailTags, error) {
	var tags EmailTags
	year, month, day := time.Now().Date()
	tags = EmailTags{
		URL:                  "https://audits.qov",
		TargetFirstName:      target.FirstName,
		TargetPosition:       target.Position,
		CustomerCity:         customer.City,
		TimeCurrentDateShort: strconv.Itoa(int(month)) + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(year%100),
		FakeFirstNameMale:    "Michael",
	}
	return tags, nil
}

func SendEmailMailgun(email MailgunEmail, client *mailgun.MailgunImpl) (string, error) {
	m := client.NewMessage(email.From, email.Subject, email.Text, email.To)
	if email.Cc != "" {
		m.AddCC(email.Html)
	}
	if email.Bcc != "" {
		m.AddBCC(email.Html)
	}
	if email.Html != "" {
		m.SetHtml(email.Html)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, id, err := client.Send(ctx, m)
	return id, err
}
