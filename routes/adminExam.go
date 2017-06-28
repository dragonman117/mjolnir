package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/session"
)

func initExamRoutes(m *macaron.Macaron){
	m.Group("/exam", func(){
		m.Get("/", examHomePage)
		m.Get("/add", examAdd)
	}, IsLoggedIn)
}

func examHomePage(ctx *macaron.Context, s session.Store){
	getGlobalValues(s, ctx.Data)
	ctx.HTML(200, "admin/adminExamHome")
}

func examAdd(ctx *macaron.Context, s session.Store){
	getGlobalValues(s, ctx.Data)
	ctx.Data["type"] = "Add"
	ctx.HTML(200, "admin/adminExamAddEdit")
}