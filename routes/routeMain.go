package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/dragonman117/oauth2"
	"github.com/go-macaron/session"
	"mjolnir/database"
)

func InitRoutes(m *macaron.Macaron){
	m.Get("/", mainPage)
	initLoginFunctions(m)
	initAdminDash(m)
}

func isLogedIn(token oauth2.Tokens) bool{
	if token.Expired() {
		return false
	}
	return true
}

func getGlobalValues(sess session.Store, data map[string]interface{}){
	cUser, ok := sess.Get("currentUser").(database.User)
	if ok{
		data["currentUser"] = cUser
	}
}

func IsLoggedIn(ctx *macaron.Context, token oauth2.Tokens){
	if !isLogedIn(token){
		ctx.Redirect("/login")
	}
	return
}