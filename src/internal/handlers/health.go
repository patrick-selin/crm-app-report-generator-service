package handlers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/utils"
)

func HealthCheckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Hello World!!", nil))
}
