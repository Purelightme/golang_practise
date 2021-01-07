package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string
	TeacherID int
	Teacher Teacher
}
