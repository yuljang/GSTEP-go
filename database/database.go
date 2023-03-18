package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gdsc-ys/21days-gin/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "host=localhost user=postgres password=21daysPassword dbname=21days port=5432 sslmode=disable TimeZone=Asia/Seoul"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.User{})
}
