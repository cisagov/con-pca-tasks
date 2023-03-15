package collections

import (
	"errors"
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

// GetCycleTargets returns all targets associated with a cycle
func GetCycleTargets(cycle_id string) ([]Target, error) {
	var t []Target

	// Return targets
	cursor, err := db.TargetsCollection.Find(db.Ctx, bson.D{{Key: "cycle_id", Value: cycle_id}})
	if err != nil {
		return t, err
	}
	if err = cursor.All(db.Ctx, &t); err != nil {
		return t, err
	}
	return t, nil
}

// GetCycleTargets returns all targets associated with a cycle
func GetCycleTargetsForProcessing(cycle_id string) ([]Target, error) {
	var t []Target
	// Get the cycle
	c, err := GetCycle(cycle_id)
	if err != nil {
		return t, err
	}

	// If cycle is active, return targets
	if c.Active {
		filter := bson.D{
			{"$and",
				bson.A{
					bson.D{{Key: "cycle_id", Value: cycle_id}},
					bson.D{{Key: "sent", Value: false}},
					bson.D{{"send_date", bson.D{{"$lte", time.Now()}}}},
				},
			},
		}
		cursor, err := db.TargetsCollection.Find(db.Ctx, filter)
		if err != nil {
			return t, err
		}
		if err = cursor.All(db.Ctx, &t); err != nil {
			return t, err
		}
		return t, nil
	} else {
		return t, errors.New("Cycle is not active")
	}
}
