// database.go
package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connection initializes the database connection using environment variables or default values
func Connection() {
	host := "localhost"
	user := "admin"
	password := "12345"
	dbname := "executiontimes_db"
	port := "5432"

	DSN := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port

	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}

// RegisterInfo registers execution time information in the database
func RegisterInfo(test string, language string, algorithm string, time int64) {
	type ExecutionTime struct {
		ID        uint   `gorm:"primaryKey"`
		Test      string `gorm:"not null"`
		Language  string `gorm:"not null"`
		Algorithm string `gorm:"not null"`
		Time      int64  `gorm:"not null"`
	}

	if err := DB.AutoMigrate(&ExecutionTime{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	tx := DB.Begin()
	if tx.Error != nil {
		log.Fatalf("Failed to begin transaction: %v", tx.Error)
	}

	if err := tx.Create(&ExecutionTime{Test: test, Language: language, Algorithm: algorithm, Time: time}).Error; err != nil {
		tx.Rollback()
		log.Fatalf("Failed to insert data into database: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}
