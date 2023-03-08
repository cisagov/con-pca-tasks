package collections

import (
	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Phish struct {
	Name    string `bson:"name"`
	Subject string `bson:"subject"`
	Html    string `bson:"html"`
	Text    string `bson:"text"`
	Retired bool   `bson:"retired"`
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
