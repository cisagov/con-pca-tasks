package notifications

import (
	"bytes"
	"log"
	"net/http"

	"github.com/cisagov/con-pca-tasks/aws"
	"github.com/cisagov/con-pca-tasks/database/collections"
)

var (
	ApiUrl string
)

// generatePDF generates a pdf file from the API
func generatePDF(cycleId, taskType string) ([]byte, error) {
	// Get the pdf file from the API
	resp, err := http.Get(ApiUrl + "/api/cycle/" + cycleId + "/reports/" + taskType + "/pdf/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	return buf.Bytes(), nil
}

// Manager manages notification emails by task type.
func Manager(cycleId, tasktype string) {
	// Get the cycle by Id
	c, err := collections.GetCycle(cycleId)
	if err != nil {
		log.Println("Get cycle error:", err.Error())
	}
	// Get the subscription by Id
	s, err := collections.GetSubscription(c.SubscriptionId)
	if err != nil {
		log.Println("Get subscription error: ", err.Error())
	}

	// Get the notification email by task type
	n, err := collections.GetNotification(tasktype + "_report")
	if err != nil {
		log.Println("Notification error: ", err.Error())
	}

	// Initialize an SES email
	email := aws.NewSESEmail()

	// Render html and text template
	tmpl := Template{FirstName: s.PrimaryContact.FirstName, LastName: s.PrimaryContact.LastName}
	textBody := tmpl.Render(n.Text)
	textHtml := tmpl.Render(n.Html)

	// Build the pdf file
	pdfFileName :=
		"CISA_PCA_" + tasktype + "_report_" + s.Name + ".pdf"
	pdfFile, err := generatePDF(cycleId, tasktype)
	if err != nil {
		log.Println("Generate PDF error: ", err.Error())
	}

	// Build and send the notification email
	email.BuildMessage(s.PrimaryContact.Email, "", s.AdminEmail, n.Subject, textHtml, textBody, pdfFileName, pdfFile)
	log.Printf("Sending email to: %s, bcc: %s", s.PrimaryContact.Email, s.AdminEmail)
	email.Send()
}
