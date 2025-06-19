package db

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ExternalID  string     `gorm:"unique;not null" json:"external_id"`
	Name        string     `gorm:"not null" json:"name"`
	Email       string     `gorm:"not null" json:"email"`
	DateOfBirth *time.Time `gorm:"not null" json:"date_of_birth"`
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("userService.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto-migrate the User model
	err = DB.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
