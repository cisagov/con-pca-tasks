package collections

import (
	"time"

	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	CustomerContact struct {
		FirstName   string `bson:"first_name"`
		LastName    string `bson:"last_name"`
		Title       string `bson:"title"`
		OfficePhone string `bson:"office_phone"`
		MobilePhone string `bson:"mobile_phone"`
		Email       string `bson:"email"`
		Notes       string `bson:"notes"`
		Active      bool   `bson:"active"`
	}

	Customer struct {
		Name                 string            `bson:"name"`
		Identifier           string            `bson:"identifier"`
		StakeholderShortname string            `bson:"stakeholder_shortname"`
		Address1             string            `bson:"address_1"`
		Address2             string            `bson:"address_2"`
		City                 string            `bson:"city"`
		State                string            `bson:"state"`
		ZipCode              string            `bson:"zip_code"`
		CustomerType         string            `bson:"customer_type"`
		ContactList          []CustomerContact `bson:"contact_list"`
		Industry             string            `bson:"industry"`
		Sector               string            `bson:"sector"`
		Domain               string            `bson:"domain"`
		AppendixADate        time.Time         `bson:"appendix_a_date"`
		Archived             bool              `bson:"archived"`
		ArchivedDescription  string            `bson:"archived_description"`
	}
)

// GetCustomer returns a notification template by task name
func GetCustomer(id string) (Customer, error) {
	var c Customer

	// Convert the string id to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c, err
	}

	err = db.CustomersCollection.
		FindOne(db.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&c)
	if err != nil {
		return c, err
	}
	return c, nil
}
