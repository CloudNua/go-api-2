package db

import (
	"log"

	"github.com/CloudNua/go-api-2/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(url string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Comment{})

	return db
}
