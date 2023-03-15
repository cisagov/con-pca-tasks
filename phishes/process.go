package phishes

import (
	"errors"
	"time"

	"github.com/cisagov/con-pca-tasks/database/collections"
	"github.com/cisagov/con-pca-tasks/services/mailgun"
	"go.mongodb.org/mongo-driver/bson"
)

func ProcessTarget(t collections.Target) (collections.Target, error) {
	if !t.Sent {
		// Get the phish template
		template, err := collections.GetPhishByID(t.TemplateID)
		if err != nil {
			return t, err
		}
		// Build the email message
		email := mailgun.NewMailgunEmail()
		email.BuildMessage(template.FromAddress, t.Email, "", "", template.Subject, template.Text, template.Html)
		if err != nil {
			return t, err
		}

		// Get the subscription
		subscription, err := collections.GetSubscription(t.SubscriptionID)
		if err != nil {
			return t, err
		}
		// Get the sending profile
		sending_profile, err := collections.GetSendingProfile(subscription.SendingProfileID)
		if err != nil {
			return t, err
		}

		// Start the mailgun client
		client := mailgun.GetMailgunClient(sending_profile.MailgunDomain, sending_profile.MailgunAPIKey)
		// Send the email
		_, sending_err := mailgun.SendEmailMailgun(*email, client)
		if sending_err != nil {
			return t, sending_err
		}

		// Update the target info
		t.Sent = true
		t.SentDate = time.Now()
		t.UpdateTarget(bson.D{{"sent", t.Sent}, {"sent_date", t.SentDate}})
		// Return the updated target
		return t, nil
	} else {
		return t, errors.New("Target phish has already been sent")
	}
}
