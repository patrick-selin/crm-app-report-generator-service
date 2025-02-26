package repository

import (
	"log"

	"github.com/patrick-selin/crm-app-report-generator-service/internal/database"
	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}


func (r *OrderRepository) GetOrdersByIDs(orderIDs []string) ([]models.Order, error) {
	var orders []models.Order

	if err := database.DB.Preload("OrderItems").
		Where("order_id IN ?", orderIDs).
		Find(&orders).Error; err != nil {
		log.Printf("OrderRepository.GetOrdersByIDs: Failed to retrieve orders: %v", err)
		return nil, err
	}

	log.Printf("OrderRepository.GetOrdersByIDs: Retrieved %d orders", len(orders))
	return orders, nil
}
