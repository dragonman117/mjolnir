package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/jinzhu/gorm"
	"github.com/go-macaron/session"
	"github.com/go-macaron/binding"
	"mjolnir/database"
	"fmt"
)

type tmpRec struct{
	Name string
}

func tstMiddle(){
	fmt.Println("tst middle ran...")
}

func initSectionRoutes(m *macaron.Macaron){
	m.Group("/sections",func(){
		m.Get("/", sectionHomePage)
		m.Get("/add", sectionAddNew)
		m.Post("/add", binding.Bind(database.Section{}), sectionAddNewPost)
		m.Get("/edit/:id", sectionEdit)
		m.Post("/edit/:id", binding.Bind(database.Section{}), sectionEditPost)
	}, IsLoggedIn)
}


func sectionHomePage(ctx *macaron.Context, db *gorm.DB, s session.Store){
	var sections []database.Section
	db.Where("archived = 0").Find(&sections)
	getGlobalValues(s, ctx.Data)
	ctx.Data["sections"] = sections
	fmt.Println(sections)
	ctx.HTML(200, "admin/adminSectionHome")
}

func sectionAddNew(ctx *macaron.Context, db *gorm.DB, s session.Store){
	getGlobalValues(s, ctx.Data)
	ctx.Data["type"] = "Add"
	ctx.HTML(200, "admin/adminSectionAddEdit")
}

func sectionAddNewPost(ctx *macaron.Context, db *gorm.DB, received database.Section, s session.Store){
	received.Archived = false;
	user, _ := s.Get("users").(database.User)
	received.UserID = user.ID
	db.Create(&received)
	ctx.Redirect("/admin/sections")
}

func sectionEdit(ctx *macaron.Context, db *gorm.DB, s session.Store){
	var current database.Section
	db.Find(&current, ctx.Params("id"))
	getGlobalValues(s, ctx.Data)
	ctx.Data["current"] = current.Name
	ctx.Data["type"] = "Edit"
	ctx.Data["edit"] = true
	ctx.HTML(200, "admin/adminSectionAddEdit")
}

func sectionEditPost(ctx *macaron.Context, db *gorm.DB, s session.Store, recieved database.Section){
	var old database.Section
	fmt.Println(recieved)
	db.First(&old, ctx.Params("id"))
	old.Name = recieved.Name
	old.Archived = recieved.Archived
	db.Save(&old)
	getGlobalValues(s, ctx.Data)
	ctx.Data["current"] = old.Name
	ctx.Data["type"] = "Edit"
	ctx.Data["edit"] = true
	ctx.HTML(200, "admin/adminSectionAddEdit")
}

