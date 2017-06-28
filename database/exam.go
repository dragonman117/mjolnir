package database

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Exam struct {
	gorm.Model
	Name string
	Rules string
	StartTime time.Time
	EndTime time.Time

	UserID uint
	Questions []Question
}