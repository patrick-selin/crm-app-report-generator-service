package repository

import (
	"log"

	"github.com/patrick-selin/crm-app-report-generator-service/internal/database"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
)

type ReportRepository struct{}

func NewReportRepository() *ReportRepository {
	return &ReportRepository{}
}

func (r *ReportRepository) StoreReportMetadata(report models.Report) error {
	if err := database.DB.Create(&report).Error; err != nil {
		log.Printf("Error saving report metadata: %v", err)
		return err
	}
	log.Printf("Report metadata stored: %v", report)
	return nil
}
