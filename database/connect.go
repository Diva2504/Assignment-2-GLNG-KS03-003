package database

import (
	"fmt"
	"os"

	"github.com/Diva2504/Assignment-2-GLNG-KS03-003/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	userName = os.Getenv("DB_USERNAME")
	dbName   = os.Getenv("DB_NAME")
	dbPass   = os.Getenv("DB_PASSWORD")
	dbPort   = os.Getenv("DB_PORT")
	dbHost   = os.Getenv("DB_HOST")
	db       *gorm.DB
	err      error
)

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, userName, dbPass, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
	db.Debug().AutoMigrate(models.Orders{}, models.Items{})
}

func GetDB() *gorm.DB {
	return db
}
