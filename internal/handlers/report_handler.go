// src/internal/handlers/lambda/report_handler.go
package lambdahandlers

import (
	"context"

	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/services"
)

type Request struct {
	ReportID string `json:"report_id"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, req Request) (Response, error) {
	// Call the same service used by Echo
	service := services.NewReportService()
	reportID, err := service.CreateReport(models.ReportRequest{
		OrderIDs: []string{req.ReportID},
		ReportType: "pdf", 
	})
	if err != nil {
		return Response{Message: "Failed to generate report"}, err
	}

	return Response{
		Message: "Report generation started with ReportID: " + reportID,
	}, nil
}
