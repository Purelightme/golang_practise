package model

import "gorm.io/gorm"

type User struct {
	ID int64
	Name string
	Age int64
	gorm.Model
	//UserProfile UserProfile
}
