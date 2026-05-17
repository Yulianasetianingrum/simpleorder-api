package domain

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	InvoiceNumber string         `gorm:"uniqueIndex;not null" json:"invoice_number"`
	CustomerID    uint           `gorm:"not null" json:"customer_id"`
	Customer      Customer       `gorm:"foreignKey:CustomerID" json:"customer"`
	TotalAmount   float64        `gorm:"not null" json:"total_amount"`
	Status        string         `gorm:"not null;default:'pending'" json:"status"`
	Items         []OrderItem    `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type OrderItem struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	OrderID   uint           `gorm:"not null" json:"order_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int            `gorm:"not null" json:"quantity" validate:"required,gt=0"`
	Price     float64        `gorm:"not null" json:"price"` // Price at the time of order
	SubTotal  float64        `gorm:"not null" json:"sub_total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateOrderRequest struct {
	CustomerID uint                   `json:"customer_id" validate:"required"`
	Items      []CreateOrderItemRequest `json:"items" validate:"required,min=1,dive"`
}

type CreateOrderItemRequest struct {
	ProductID uint `json:"product_id" validate:"required"`
	Quantity  int  `json:"quantity" validate:"required,gt=0"`
}
