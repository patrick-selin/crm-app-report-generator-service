// lambda/main.go
package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/handlers"
)

type Request struct {
	ReportID string `json:"report_id"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, req Request) (Response, error) {
	message := handlers.GenerateReport(req.ReportID)
	return Response{
		Message: message,
	}, nil
}

func main() {
	lambda.Start(handler)
}
