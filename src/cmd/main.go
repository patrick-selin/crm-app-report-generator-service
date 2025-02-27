package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/database"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/handlers"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/utils"
)

func main() {
	database.ConnectDB()
	e := echo.New()
	e.Validator = utils.NewValidator()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	

	e.GET("/health", handlers.HealthCheckHandler)

	api := e.Group("/api/v1/reports")
	
	api.POST("/new", handlers.CreateReportHandler) 
	api.GET("/orders", handlers.GetAllOrdersHandler)

	port := "8080"
	log.Printf("Starting server on port %s", port)

	// DEBUG:
	for _, route := range e.Routes() {
		fmt.Printf("Method: %s | Path: %s\n", route.Method, route.Path)
	}

	e.Logger.Fatal(e.Start(":" + port))
}
