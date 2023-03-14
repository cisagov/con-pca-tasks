package mailgun

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

// BuildMessage builds email context.
func (e *MailgunEmail) BuildMessage(to, cc, bcc, subject, html, text string) {
	e.To = to
	e.Cc = cc
	e.Bcc = bcc
	e.Subject = subject
	e.Html = html
	e.Text = text
}

// Send sends the email using Mailgun.
func (e *MailgunEmail) Send() error {
	return nil
}
