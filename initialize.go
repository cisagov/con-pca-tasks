package main

import (
	db "github.com/cisagov/con-pca-tasks/database"
	"github.com/cisagov/con-pca-tasks/services/aws"
)

// Initialize the database and AWS clients
func Init() {
	// Connect to the database
	db.InitDB()

	// Initialize SES email
	aws.SESEmailClient()
}
