package repository

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/storage"
)

type DynamoDBRepository struct {
	Client *dynamodb.Client
	Table  string
}

func NewDynamoDBRepository() *DynamoDBRepository {
	return &DynamoDBRepository{
		Client: storage.NewDynamoDBClient(),
		Table:  "ReportMetadataTable", // muista vaihtaa
	}
}

func (r *DynamoDBRepository) StoreReportMetadata(metadata models.ReportMetadata) error {
	metadata.CreatedAt = time.Now().UTC()
	item, err := attributevalue.MarshalMap(metadata)
	if err != nil {
		log.Printf("DynamoDBRepository: Failed to marshal metadata: %v", err)
		return err
	}

	_, err = r.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(r.Table),
		Item:      item,
	})
	if err != nil {
		log.Printf("DynamoDBRepository: Failed to store metadata: %v", err)
		return err
	}

	log.Printf("DynamoDBRepository: Report metadata stored: %v", metadata)
	return nil
}
