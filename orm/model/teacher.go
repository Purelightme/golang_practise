package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name string `gorm:"size:256"`
	Age uint8 `gorm:"not_null;default:0"`
}