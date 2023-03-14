package collections

import (
	"time"

	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	TargetEmail struct {
		Email     string `bson:"email"`
		FirstName string `bson:"first_name"`
		LastName  string `bson:"last_name"`
		Position  string `bson:"position"`
	}

	PrimaryContact struct {
		FirstName   string `bson:"first_name"`
		LastName    string `bson:"last_name"`
		Title       string `bson:"title"`
		OfficePhone string `bson:"office_phone"`
		MobilePhone string `bson:"mobile_phone"`
		Email       string `bson:"email"`
		Notes       string `bson:"notes"`
		Active      bool   `bson:"active"`
	}

	SubscriptionTasks struct {
		TaskUUID      string    `bson:"task_uuid"`
		TaskType      string    `bson:"task_type"`
		ScheduledDate time.Time `bson:"scheduled_date"`
		Executed      bool      `bson:"executed"`
		ExecutedDate  time.Time `bson:"executed_date"`
		Error         string    `bson:"error"`
	}

	Subscription struct {
		Name                   string              `bson:"name"`
		CustomerID             string              `bson:"customer_id"`
		SendingProfileID       string              `bson:"sending_profile_id"`
		TargetDomain           string              `bson:"target_domain"`
		Customer               string              `bson:"customer"`
		StartDate              time.Time           `bson:"start_date"`
		PrimaryContact         PrimaryContact      `bson:"primary_contact"`
		AdminEmail             string              `bson:"admin_email"`
		OperatorEmail          string              `bson:"operator_email"`
		Status                 string              `bson:"status"`
		CycleStartDate         string              `bson:"cycle_start_date"`
		TargetEmailList        []TargetEmail       `bson:"target_email_list"`
		TemplatesSelected      []string            `bson:"templates_selected"`
		NextTemplates          []string            `bson:"next_templates"`
		ContinuousSubscription bool                `bson:"continuous_subscription"`
		BufferTimeMinutes      int                 `bson:"buffer_time_minutes"`
		CycleLengthMinutes     int                 `bson:"cycle_length_minutes"`
		CooldownMinutes        int                 `bson:"cooldown_minutes"`
		ReportFrequencyMinutes int                 `bson:"report_frequency_minutes"`
		Tasks                  []SubscriptionTasks `bson:"tasks"`
		Processing             bool                `bson:"processing"`
		Archived               bool                `bson:"archived"`
	}
)

// GetSubscription returns a notification template by task name
func GetSubscription(id string) (Subscription, error) {
	var s Subscription

	// Convert the string id to an ObjectID
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return s, err
	}

	err = db.SubscriptionsCollection.
		FindOne(db.Ctx, bson.D{{Key: "_id", Value: objectId}}).
		Decode(&s)
	if err != nil {
		return s, err
	}
	return s, nil
}
