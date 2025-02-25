package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/handlers"
)

func main() {
	e := echo.New()

	e.GET("/health", handlers.HealthCheckHandler)

	port := "8080"
	log.Printf("Starting server on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}