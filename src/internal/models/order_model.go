package models

import (
	"time"
)

type Order struct {
	OrderID     string      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"order_id"`
	CustomerID  string      `gorm:"type:uuid;not null" json:"customer_id"`
	TotalAmount float64     `gorm:"type:numeric(10,2);not null" json:"total_amount"`
	OrderStatus string      `gorm:"type:order_status;not null;default:'Pending'" json:"order_status"`
	OrderDate   time.Time   `gorm:"type:timestamp without time zone;default:now()" json:"order_date"`
	CreatedAt   time.Time   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime" json:"updated_at"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnDelete:CASCADE;" json:"order_items"`
}

type OrderItem struct {
	OrderItemID string  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"order_item_id"`
	OrderID     string  `gorm:"type:uuid;not null" json:"order_id"`
	ProductID   string  `gorm:"type:uuid;not null" json:"product_id"`
	Quantity    int     `gorm:"type:int;not null" json:"quantity"`
	Price       float64 `gorm:"type:numeric(10,2);not null" json:"price"`
}
