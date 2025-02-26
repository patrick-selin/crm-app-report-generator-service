package handlers

import (
	"log"
	"os"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/repository"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/services"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/utils"
)

func GetAllOrdersHandler(c echo.Context) error {
	// Ensure logs are printed to stdout
	log.SetOutput(os.Stdout)

	log.Println("GetAllOrdersHandler: Start")

	service := services.NewOrderService(repository.NewOrderRepository())
	orders, err := service.GetAllOrders()
	if err != nil {
		log.Printf("GetAllOrdersHandler: Error: %v", err)
		return c.JSON(http.StatusInternalServerError, utils.NewErrorResponse(err.Error()))
	}

	log.Printf("GetAllOrdersHandler: Success. Orders: %v", orders)
	return c.JSON(http.StatusOK, utils.NewSuccessResponse("Orders retrieved successfully", orders))
}
