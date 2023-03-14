package collections

import (
	db "github.com/cisagov/con-pca-tasks/database"
	"go.mongodb.org/mongo-driver/bson"
)

type Notification struct {
	Name          string `bson:"name"`
	Subject       string `bson:"subject"`
	Html          string `bson:"html"`
	TaskName      string `bson:"task_name"`
	Text          string `bson:"text"`
	HasAttachment bool   `bson:"has_attachment"`
}

// GetNotification returns a notification template by task name
func GetNotification(TaskName string) (Notification, error) {
	var n Notification
	err := db.NotificationsCollection.
		FindOne(db.Ctx, bson.D{{Key: "task_name", Value: TaskName}}).
		Decode(&n)
	if err != nil {
		return n, err
	}
	return n, nil
}
