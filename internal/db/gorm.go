package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORM database configuration

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=root password=1234 dbname=dev_cart port=5431 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("db connection failed")
	}
	log.Println("✅ database connected")

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("failed to get database instance")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Start background connection health check in goroutine
	// go healthCheckConnections()
}

// Health check connections in background goroutine
// func healthCheckConnections() {
// 	ticker := time.NewTicker(60 * time.Second)
// 	defer ticker.Stop()

// 	for range ticker.C {
// 		go func() {
// 			if err := DB.Raw("SELECT 1").Error; err != nil {
// 				log.Printf("⚠️  Database health check failed: %v", err)
// 			} else {
// 				log.Println("✅ Database health check passed")
// 			}
// 		}()
// 	}
// }
