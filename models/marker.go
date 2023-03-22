package models

import (
	"gorm.io/gorm"
)

type Marker struct {
	gorm.Model
	Longitude float64
	Latitude  float64
	Message   string
	Address   string
	IconID    int
	UserID    int
}

// maybe Longitude + Latitude is better candidate for PK
