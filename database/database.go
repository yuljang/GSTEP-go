package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/yuljang/GSTEP-go/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := os.Getenv("GSTEP_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Marker{})
	DB.AutoMigrate(&models.Mission{})
	DB.AutoMigrate(&models.User{})
}
