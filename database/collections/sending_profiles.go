package collections

import (
	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Header struct {
		key   string `bson:"key"`
		value string `bson:"value"`
	}

	SendingProfile struct {
		Name          string   `bson:"name"`
		InterfaceType string   `bson:"interface_type"`
		FromAddress   string   `bson:"from_address"`
		Headers       []Header `bson:"headers"`
		SendingIPs    string   `bson:"sending_ips"`
		SMTPUsername  string   `bson:"smtp_username"`
		SMTPPassword  string   `bson:"smtp_password"`
		SMTPHost      string   `bson:"smtp_host"`
		MailgunDomain string   `bson:"mailgun_domain"`
		MailgunAPIKey string   `bson:"mailgun_api_key"`
		SESRoleArn    string   `bson:"ses_role_arn"`
	}
)

// GetSendingProfile returns a sending profile by id string
func GetSendingProfile(id string) (SendingProfile, error) {
	var sp SendingProfile

	// Convert the string id to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return sp, err
	}

	err = db.SendingProfilesCollection.
		FindOne(db.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&sp)
	if err != nil {
		return sp, err
	}
	return sp, nil
}
