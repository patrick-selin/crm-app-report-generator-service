// internal/storage/dynamodb.go
package storage

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("NewDynamoDBClient: Failed to load AWS config: %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)
	return client
}
