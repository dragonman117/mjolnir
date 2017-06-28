package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/oauth2"
	"github.com/go-macaron/session"
	"github.com/jinzhu/gorm"
	"mjolnir/database"
)

func mainPage(ctx *macaron.Context, token oauth2.Tokens, s session.Store, db *gorm.DB){
	if !isLogedIn(token){
		ctx.Redirect("/login")
		return
	}
	testLogin(ctx, token, s, db)
	_, ok := s.Get("currentUser").(database.User)
	if(!ok){
		var cUser database.User
		suser, ok := s.Get("user").(map[string]interface{})
		if ok {
			db.Where("auth_id = ?", suser["user_id"].(string)).First(&cUser)
			if cUser.ID != 0 {
				s.Set("currentUser", cUser)
			}
		}
	}
	getGlobalValues(s, ctx.Data)
	ctx.Data["name"] = "Awsome"
	ctx.HTML(200, "mainDashboard")
}

func testLogin(ctx *macaron.Context, tokens oauth2.Tokens, s session.Store, db *gorm.DB){
	if tokens.Expired() {
		ctx.Redirect("/login")
	}
	if isFirstLogin(s, db){
		var num int
		db.Table("users").Count(&num)
		if num == 0{
			ctx.Redirect("/install/admin")
		}else{
			ctx.Redirect("/addUser")
		}
	}
}