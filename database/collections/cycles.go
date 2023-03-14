package collections

import (
	"time"

	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cycle struct {
	SubscriptionId string    `bson:"subscription_id"`
	TemplateIds    []string  `bson:"template_ids"`
	StartDate      time.Time `bson:"start_date"`
	EndDate        time.Time `bson:"end_date"`
	SendByDate     time.Time `bson:"send_by_date"`
	Active         bool      `bson:"active"`
	TargetCount    int       `bson:"target_count"`
}

// GetCycle returns a cycle by id
func GetCycle(id string) (Cycle, error) {
	var c Cycle

	// Convert the string id to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c, err
	}

	err = db.CyclesCollection.
		FindOne(db.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}
