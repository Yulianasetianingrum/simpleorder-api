package database

import (
	"fmt"
	"log"
	"simpleorder/internal/config"
	"simpleorder/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully")

	// Auto Migrate
	err = db.AutoMigrate(
		&domain.User{},
		&domain.Customer{},
		&domain.Product{},
		&domain.Order{},
		&domain.OrderItem{},
	)
	if err != nil {
		return nil, err
	}
	
	log.Println("Database migration completed")

	return db, nil
}
