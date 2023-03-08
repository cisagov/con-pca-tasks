package notifications

import (
	"log"

	"github.com/cisagov/con-pca-tasks/aws"
	"github.com/cisagov/con-pca-tasks/database/collections"
)

// Manager manages notification emails by task type.
func Manager(CycleId, tasktype string) {
	// Get the cycle by Id
	c, err := collections.GetCycle(CycleId)
	if err != nil {
		log.Println("Get cycle error:", err.Error())
	}
	// Get the subscription by Id
	s, err := collections.GetSubscription(c.SubscriptionId)
	if err != nil {
		log.Println("Get subscription error: ", err.Error())
	}

	// Get the notification email by task type
	n, err := collections.GetNotification(tasktype)
	if err != nil {
		log.Println("Notification error: ", err.Error())
	}

	// Initialize an SES email
	email := aws.NewSESEmail()

	// Render the template
	tmpl := Template{FirstName: s.PrimaryContact.FirstName, LastName: s.PrimaryContact.LastName}
	textBody := tmpl.Render(n.Text)

	// Build and send the notification email
	email.BuildMessage(s.PrimaryContact.Email, "", s.AdminEmail, n.Subject, textBody)
	log.Printf("Sending email to: %s, bcc: %s", s.PrimaryContact.Email, s.AdminEmail)
	email.Send()
}
