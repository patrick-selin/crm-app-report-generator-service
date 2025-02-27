package storage

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
)

type ReportStorage struct {
	S3Client      *s3.Client
	DynamoClient  *dynamodb.Client
	Bucket        string
	DynamoTable   string
}

// Initializes the storage with S3 and DynamoDB clients
func NewReportStorage() *ReportStorage {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("ReportStorage: Failed to load AWS config: %v", err)
	}

	return &ReportStorage{
		S3Client:     s3.NewFromConfig(cfg),
		DynamoClient: dynamodb.NewFromConfig(cfg),
		Bucket:       "report-generator-bucket",
		DynamoTable:  "ReportMetadataTable",   
	}
}

// UploadFile uploads a file to S3 and returns its URL
func (r *ReportStorage) UploadFile(filename string, data []byte) (string, error) {
	uploader := manager.NewUploader(r.S3Client)

	key := fmt.Sprintf("reports/%s/%s", uuid.New().String(), filename)
	log.Printf("ReportStorage.UploadFile: Uploading file with key: %s", key)

	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(r.Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("application/octet-stream"),
		ACL:         types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		log.Printf("ReportStorage.UploadFile: Failed to upload file: %v", err)
		return "", err
	}

	url, err := r.GetPresignedURL(key)
	if err != nil {
		log.Printf("ReportStorage.UploadFile: Failed to get pre-signed URL: %v", err)
		return "", err
	}

	log.Printf("ReportStorage.UploadFile: File uploaded successfully with URL: %s", url)
	return url, nil
}

// Generate a pre-signed URL for an S3 object
func (r *ReportStorage) GetPresignedURL(key string) (string, error) {
	presignClient := s3.NewPresignClient(r.S3Client)

	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(r.Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		log.Printf("ReportStorage.GetPresignedURL: Failed to create presigned URL: %v", err)
		return "", err
	}

	return req.URL, nil
}

func (r *ReportStorage) StoreReportMetadata(reportMetadata models.ReportMetadata) error {
	// Muista: Implement DynamoDB put operation
	log.Printf("ReportStorage.StoreReportMetadata: Saving report metadata: %+v", reportMetadata)
	return nil
}
