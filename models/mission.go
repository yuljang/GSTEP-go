package models

import (
	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	Title   string
	Content string
	IconID  uint
}
