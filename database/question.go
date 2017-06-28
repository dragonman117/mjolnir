package database

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Name string
	Prompt string
	Order int
	ExamID uint
	TestCases []TestCase
}
