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
	IconID    int32
}

// maybe Longitude + Latitude is better candidate for PK
