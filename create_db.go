package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Connect to the default 'postgres' database
	dsn := "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to default postgres db: %v", err)
	}

	// Create simpleorder database
	err = db.Exec("CREATE DATABASE simpleorder").Error
	if err != nil {
		log.Fatalf("Failed to create simpleorder database: %v", err)
	}

	fmt.Println("Database 'simpleorder' created successfully!")
}
