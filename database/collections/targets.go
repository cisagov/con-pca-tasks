package collections

import (
	"time"

	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	TimelineDetails struct {
		UserAgent string `bson:"user_agent"`
		Ip        string `bson:"ip"`
		AsnOrg    string `bson:"asn_org"`
		City      string `bson:"city"`
		Country   string `bson:"country"`
	}

	TargetTimeline struct {
		Time    time.Time       `bson:"time"`
		Message string          `bson:"message"`
		Details TimelineDetails `bson:"details"`
	}

	Target struct {
		ID                primitive.ObjectID `bson:"_id"`
		CycleID           string             `bson:"cycle_id"`
		SubscriptionID    string             `bson:"subscription_id"`
		TemplateID        string             `bson:"template_id"`
		Email             string             `bson:"email"`
		FirstName         string             `bson:"first_name"`
		LastName          string             `bson:"last_name"`
		Position          string             `bson:"position"`
		DeceptionLevel    string             `bson:"deception_level"`
		DeceptionLevelInt int                `bson:"deception_level_int"`
		SendDate          time.Time          `bson:"send_date"`
		Sent              bool               `bson:"sent"`
		SentDate          time.Time          `bson:"sent_date"`
		Error             string             `bson:"error"`
		Timeline          TargetTimeline     `bson:"timeline"`
	}
)

// GetTarget returns a phishing target by id
func GetTarget(id string) (Target, error) {
	var t Target

	// Convert the string id to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return t, err
	}

	err = db.TargetsCollection.
		FindOne(db.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}

// UpdateTarget updates a phishing target with new data
func (t *Target) UpdateTarget(data bson.D) (*mongo.UpdateResult, error) {
	result, err := db.TargetsCollection.UpdateOne(db.Ctx, bson.D{{Key: "_id", Value: t.ID}}, bson.D{{"$set", data}})
	if err != nil {
		return result, err
	}
	return result, err
}
