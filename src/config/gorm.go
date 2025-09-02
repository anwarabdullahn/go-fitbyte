package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewGorm creates a new GORM database connection using PostgreSQL
func NewGorm(config *viper.Viper) *gorm.DB {
	username := config.GetString("database.username")
	password := config.GetString("database.password")
	host := config.GetString("database.host")
	port := config.GetInt("database.ports")
	dbName := config.GetString("database.name")
	maxConnections := config.GetInt("database.maxConnections")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		host, username, password, dbName, port)

	if username == "" || password == "" || host == "" || dbName == "" {
		log.Fatal("Database credentials are required")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Get underlying sql.DB to configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get underlying sql.DB:", err)
	}

	// Configure connection pool
	if maxConnections > 0 {
		sqlDB.SetMaxOpenConns(maxConnections)
	} else {
		sqlDB.SetMaxOpenConns(10) // default
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Test connection
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully")
	return db
}
