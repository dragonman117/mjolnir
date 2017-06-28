package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db gorm.DB

func InitDatabase() *gorm.DB{
	db, err := gorm.Open("sqlite3", "dev.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&Section{}, &Exam{}, &Question{}, &TestCase{}, &User{})

	//Todo: On move to full sql need to enable these
	//db.Model(&Exam{}).AddForeignKey("section_id", "sections(id)", "RESTRICT", "RESTRICT")
	//db.Model(&Question{}).AddForeignKey("exam_id", "exams(id)", "RESTRICT", "RESTRICT")
	//db.Model(&TestCase{}).AddForeignKey("question_id", "questions(id)", "RESTRICT", "RESTRICT")

	return db
}

func CloseDatabase(){
	db.Close();
}