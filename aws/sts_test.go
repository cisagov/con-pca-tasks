package aws

import (
	"testing"
)

// TestAssumedRoleUser tests the AssumedRoleUser function
func TestAssumedRoleUser(t *testing.T) {
	// Redeclare the fromRoleArn variable to avoid using the environment variable
	fromRoleArn = "arn:aws:iam::123456789012:role/role-name"

	cfg := AssumedRoleUser()

	if cfg.Region != "us-east-1" {
		t.Errorf("Region should be us-east-1, but got %s", cfg.Region)
	}

	if cfg.Credentials == nil {
		t.Error("Credentials should not be nil")
	}
}
