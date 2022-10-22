package config

import (
	"MyGram/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB(configuration Config) {
	db, err = gorm.Open(postgres.Open(configuration.Get("DB_CONFIG")), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	db.AutoMigrate(entity.User{}, entity.Photo{}, entity.Comment{}, entity.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
