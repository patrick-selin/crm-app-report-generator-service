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
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
}

func NewS3Storage() *S3Storage {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-north-1"))
	if err != nil {
		log.Fatalf("S3Storage: Failed to load AWS config: %v", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Storage{
		Client: client,
		Bucket: "report-generator-bucket",
	}
}

func (s *S3Storage) UploadFile(filename string, data []byte) (string, error) {
	uploader := manager.NewUploader(s.Client)

	key := fmt.Sprintf("reports/%s/%s", uuid.New().String(), filename)
	log.Printf("S3Storage.UploadFile: Uploading file with key: %s", key)

	_, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
		ContentType: aws.String("application/octet-stream"),
		ACL:    types.ObjectCannedACLPublicRead,
	})

	if err != nil {
		log.Printf("S3Storage.UploadFile: Failed to upload file: %v", err)
		return "", err
	}

	// Generate pre-signed URL
	url, err := s.GetPresignedURL(key)
	if err != nil {
		log.Printf("S3Storage.UploadFile: Failed to get pre-signed URL: %v", err)
		return "", err
	}

	log.Printf("S3Storage.UploadFile: File uploaded successfully with URL: %s", url)
	return url, nil
}

func (s *S3Storage) GetPresignedURL(key string) (string, error) {
	presignClient := s3.NewPresignClient(s.Client)

	req, err := presignClient.PresignGetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(15*time.Minute))

	if err != nil {
		log.Printf("S3Storage.GetPresignedURL: Failed to create presigned URL: %v", err)
		return "", err
	}

	return req.URL, nil
}
