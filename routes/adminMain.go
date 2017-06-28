package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/dragonman117/oauth2"
	"github.com/go-macaron/session"
)

func initAdminDash(m *macaron.Macaron){
	m.Group("/admin", func() {
		m.Get("/dashboard", adminMain)
		initSectionRoutes(m)
		initExamRoutes(m)
	})
}

func adminMain(ctx *macaron.Context, s session.Store, token oauth2.Tokens){
	if !isLogedIn(token){
		ctx.Redirect("/login")
		return
	}
	getGlobalValues(s, ctx.Data)
	ctx.HTML(200, "admin/adminDashboard")
}