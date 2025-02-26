package services

import (
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/repository"
)

type ReportService struct {
	repo *repository.ReportRepository
}

func NewReportService() *ReportService {
	return &ReportService{
		repo: repository.NewReportRepository(),
	}
}

func (s *ReportService) CreateReport(request models.ReportRequest) (string, error) {
	reportID := uuid.New().String()

	// Mocked Report Generation (CSV/PDF/Both)
	fileURL := fmt.Sprintf("https://s3.amazonaws.com/fake-bucket/%s.%s", reportID, request.ReportType)
	report := models.Report{
		ReportID:     reportID,
		FileURL:      fileURL,
		Status:       "Pending",
		ReportType:   request.ReportType,
		IncludeItems: request.IncludeItems,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Store report metadata
	err := s.repo.StoreReportMetadata(report)
	if err != nil {
		log.Printf("Failed to store report metadata: %v", err)
		return "", err
	}

	return reportID, nil
}
