package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/patrick-selin/crm-app-report-generator-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB() {
	// Construct the DSN
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get the underlying sql.DB to manage connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v", err)
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	// Check and create ENUM type if it doesn't exist
	enumCheckQuery := `
		SELECT EXISTS (
			SELECT 1 
			FROM pg_type t 
			JOIN pg_enum e ON t.oid = e.enumtypid 
			WHERE t.typname = 'order_status'
		);
	`

	var exists bool
	err = db.Raw(enumCheckQuery).Scan(&exists).Error
	if err != nil {
		log.Fatalf("Failed to check ENUM type: %v", err)
	}

	if !exists {
		log.Println("Creating ENUM type order_status...")
		createEnumQuery := `
			CREATE TYPE order_status AS ENUM ('Pending', 'Processing', 'Shipped', 'Delivered', 'Cancelled');
		`
		err = db.Exec(createEnumQuery).Error
		if err != nil {
			log.Fatalf("Failed to create ENUM type: %v", err)
		}
		log.Println("ENUM type order_status created successfully.")
	} else {
		log.Println("ENUM type order_status already exists.")
	}


	// Run migrations
	err = db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Connected to database successfully and migrations ran!")
	DB = db
}
