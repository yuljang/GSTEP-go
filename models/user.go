package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string
	Nickname  string
	Image     string
	Country   string
	Step      int
	Milestone datatypes.JSON
	Point     int
	Progress  datatypes.JSON // -2: not discovered, 0: discovered, 1~: countable mission, count number
	Marker    Marker
}
