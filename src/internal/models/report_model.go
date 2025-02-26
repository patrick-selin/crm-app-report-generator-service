package models

import (
	"time"
)

type ReportStatus string

const (
	StatusPending   ReportStatus = "Pending"
	StatusCompleted ReportStatus = "Completed"
	StatusFailed    ReportStatus = "Failed"
)

type ReportMetadata struct {
	ReportID     string       `json:"report_id" dynamodbav:"ReportID"`
	UserID       string       `json:"user_id" dynamodbav:"UserID"`
	OrderIDs     []string     `json:"order_ids" dynamodbav:"OrderIDs"`
	ReportType   string       `json:"report_type" dynamodbav:"ReportType"`
	IncludeItems bool         `json:"include_items" dynamodbav:"IncludeItems"`
	Status       ReportStatus `json:"status" dynamodbav:"Status"`
	S3Key        string       `json:"s3_key" dynamodbav:"S3Key"`
	CreatedAt    time.Time    `json:"created_at" dynamodbav:"CreatedAt"`
	UpdatedAt    time.Time    `json:"updated_at" dynamodbav:"UpdatedAt"`
}