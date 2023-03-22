package models

import (
	"gorm.io/gorm"
)

type Marker struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Message   string
	Address   string
	IconID    int
	UserID    int
}

// maybe Latitude + Longitude is better candidate for PK
