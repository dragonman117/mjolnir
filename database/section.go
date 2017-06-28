package database

import "github.com/jinzhu/gorm"

type Section struct {
	gorm.Model
	Name string
	Archived bool

	UserID uint
	Exams []Exam
}