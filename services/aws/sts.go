package aws

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

var (
	stsClient   *sts.Client
	fromRoleArn = os.Getenv("SES_ASSUME_ROLE_ARN")
)

// AssumedRoleConfig returns an AWS configuration with assumed user role credentials
func AssumedRoleConfig() aws.Config {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Println("STS configuration error: ", err.Error())
	}

	stsClient = sts.NewFromConfig(cfg)
	provider := stscreds.NewAssumeRoleProvider(stsClient, fromRoleArn)

	cfg.Credentials = aws.NewCredentialsCache(provider)

	return cfg
}
