package collections

import (
	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Phish struct {
	Name        string `bson:"name"`
	Subject     string `bson:"subject"`
	FromAddress string `bson:"from_address"`
	Html        string `bson:"html"`
	Text        string `bson:"text"`
	Retired     bool   `bson:"retired"`
}

// GetPhish returns a phish template by name
func GetPhish(Name string) (Phish, error) {
	var p Phish
	err := db.PhishesCollection.
		FindOne(db.Ctx, bson.D{{Key: "name", Value: Name}, {Key: "retired", Value: false}}).
		Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

// GetPhish returns a phish template by id
func GetPhishByID(id string) (Phish, error) {
	var p Phish

	// Convert the string id to an ObjectID
	objectId, convert_err := primitive.ObjectIDFromHex(id)
	if convert_err != nil {
		return p, convert_err
	}

	err := db.PhishesCollection.
		FindOne(db.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}
