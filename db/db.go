package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"technodom_test/models"
)

var DB *gorm.DB

func ConnectToDataBase() {
	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=postgres password=123 dbname=links port=5432"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database...")
	}
}

func SyncDB() {
	DB.AutoMigrate(&models.Link{})
}
