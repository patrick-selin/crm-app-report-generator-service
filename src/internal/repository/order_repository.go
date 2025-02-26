package repository

import (
	"github.com/patrick-selin/crm-app-report-generator-service/internal/database"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (r *OrderRepository) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	err := database.DB.Preload("OrderItems").Order("order_date DESC").Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
