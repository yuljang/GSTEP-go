package models

import (
	"gorm.io/gorm"
)

type Mission struct {
	gorm.Model
	Type         string
	Title        string
	Content      string
	ShortContent string
	IconID       uint
}
