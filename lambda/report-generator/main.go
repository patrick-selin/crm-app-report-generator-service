// lambda/report-generator/main.go
package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/handlers"
)

func main() {
	lambda.Start(lambdahandlers.HandleRequest)
}
