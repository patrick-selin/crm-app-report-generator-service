// src/internal/handlers/lambda/report_handler.go
package handlers

import (
	"context"

	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/services"
)

type Request struct {
	OrderIDs     []string `json:"order_ids"`
	ReportType   string   `json:"report_type"`
	IncludeItems bool     `json:"include_items"`
}

type Response struct {
	Message string `json:"message"`
}

func HandleRequest(ctx context.Context, req Request) (Response, error) {
	service := services.NewReportService()

	reportID, err := service.CreateReport(models.ReportRequest{
		OrderIDs:     req.OrderIDs,
		ReportType:   req.ReportType,
		IncludeItems: req.IncludeItems,
	})
	if err != nil {
		return Response{Message: "Failed to generate report"}, err
	}

	return Response{
		Message: "Report generation started with ReportID: " + reportID,
	}, nil
}
