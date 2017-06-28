package database

import "github.com/jinzhu/gorm"

type User struct{
	gorm.Model
	AuthId string
	Email string
	FirstName string
	LastName string
	ANumber int
	IsGlobalAdmin bool
}