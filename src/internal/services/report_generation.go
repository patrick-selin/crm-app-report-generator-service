package services

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"time"

	"github.com/jung-kurt/gofpdf"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
)

// Generate CSV
func generateCSV(orders []models.Order) ([]byte, string, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	// Write header
	writer.Write([]string{"Order ID", "Customer ID", "Total Amount", "Status", "Order Date"})

	// Write order rows
	for _, order := range orders {
		row := []string{
			order.OrderID,
			order.CustomerID,
			fmt.Sprintf("%.2f", order.TotalAmount),
			order.OrderStatus,
			order.OrderDate.Format(time.RFC3339),
		}
		writer.Write(row)
	}

	writer.Flush()

	filename := fmt.Sprintf("report_%d.csv", time.Now().Unix())
	return buf.Bytes(), filename, nil
}

// Generate PDF
func generatePDF(orders []models.Order) ([]byte, string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	// Header
	pdf.Cell(40, 10, "Order Report")
	pdf.Ln(12)

	// Table Headers
	pdf.SetFont("Arial", "B", 10)
	pdf.Cell(30, 10, "Order ID")
	pdf.Cell(30, 10, "Customer ID")
	pdf.Cell(30, 10, "Amount")
	pdf.Cell(30, 10, "Status")
	pdf.Cell(30, 10, "Order Date")
	pdf.Ln(10)

	// Table Rows
	pdf.SetFont("Arial", "", 10)
	for _, order := range orders {
		pdf.Cell(30, 10, order.OrderID)
		pdf.Cell(30, 10, order.CustomerID)
		pdf.Cell(30, 10, fmt.Sprintf("%.2f", order.TotalAmount))
		pdf.Cell(30, 10, order.OrderStatus)
		pdf.Cell(30, 10, order.OrderDate.Format("2006-01-02"))
		pdf.Ln(10)
	}

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		log.Printf("generatePDF: Failed to generate PDF: %v", err)
		return nil, "", err
	}

	filename := fmt.Sprintf("report_%d.pdf", time.Now().Unix())
	return buf.Bytes(), filename, nil
}
