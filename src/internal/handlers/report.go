// handler
package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/services"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/utils"
)

// POST /api/v1/reports/new
func CreateReportHandler(c echo.Context) error {
	var request models.ReportRequest

	log.Println("Request Body:", c.Request().Body)
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Invalid request payload", err))
	}

	// Debug:
	log.Printf("Parsed Request Payload: %+v", request)

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.NewErrorResponse("Validation failed", err))
	}

	reportID, err := services.NewReportService().CreateReport(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse("Failed to create report", err))
	}

	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Report generation started", map[string]interface{}{
		"report_id": reportID,
	}))
}
