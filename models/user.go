package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Nickname string
	Image    string
	Country  string
	Level    uint
	Marker   Marker
	Mission  datatypes.JSON // -2: not discovered, 0: discovered, 1~: countable mission, count number
}
