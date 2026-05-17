package domain

import (
	"time"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name" validate:"required"`
	Description string         `json:"description"`
	Price       float64        `gorm:"not null" json:"price" validate:"required,gt=0"`
	Stock       int            `gorm:"not null;default:0" json:"stock" validate:"min=0"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
