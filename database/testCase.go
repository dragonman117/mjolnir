package database

import "github.com/jinzhu/gorm"

type TestCase struct {
	gorm.Model
	Input string
	Output string
	Public bool
	QuestionID uint
}