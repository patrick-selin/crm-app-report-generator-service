package services

import (
	"fmt"
	"log"
	"time"
	"github.com/google/uuid"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/repository"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/storage"
)

type ReportService struct {
	OrderRepo *repository.OrderRepository
	S3Storage *storage.S3Storage
}

func NewReportService() *ReportService {
	return &ReportService{
		OrderRepo: repository.NewOrderRepository(),
		S3Storage: storage.NewS3Storage(),
	}
}

func (s *ReportService) CreateReport(request models.ReportRequest) (string, error) {
	reportID := uuid.New().String()

	orders, err := s.OrderRepo.GetOrdersByIDs(request.OrderIDs)
	if err != nil {
		log.Printf("CreateReport: Failed to retrieve orders: %v", err)
		return "", err
	}

	if len(orders) == 0 {
		return "", fmt.Errorf("No orders found for the provided order IDs")
	}

	// Generate CSV and/or PDF
	var csvURL, pdfURL string
	if request.ReportType == "csv" || request.ReportType == "both" {
		csvData, csvFilename, err := generateCSV(orders)
		if err != nil {
			return "", err
		}
		csvURL, err = s.S3Storage.UploadFile(csvFilename, csvData)
		if err != nil {
			return "", err
		}
	}

	if request.ReportType == "pdf" || request.ReportType == "both" {
		pdfData, pdfFilename, err := generatePDF(orders)
		if err != nil {
			return "", err
		}
		pdfURL, err = s.S3Storage.UploadFile(pdfFilename, pdfData)
		if err != nil {
			return "", err
		}
	}

	// Store metadata in DynamoDB
	reportMetadata := models.ReportMetadata{
		ReportID:     reportID,
		UserID:       "sample-user-id", // To be retrieved from JWT claims
		OrderIDs:     request.OrderIDs,
		ReportType:   request.ReportType,
		IncludeItems: request.IncludeItems,
		Status:       models.StatusCompleted,
		S3Key:        csvURL + ", " + pdfURL,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err = s.S3Storage.StoreReportMetadata(reportMetadata)
	if err != nil {
		return "", err
	}

	log.Printf("CreateReport: Report generated successfully with ID: %s", reportID)
	return reportID, nil
}
