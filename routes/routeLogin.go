package routes

import (
	"gopkg.in/macaron.v1"
	"github.com/go-macaron/session"
	"mjolnir/database"
	"github.com/jinzhu/gorm"
	"github.com/dragonman117/oauth2"
	"regexp"
	"strconv"
	"github.com/go-macaron/binding"
)

type userAdd struct{
	Fname string
	Lname string
	Anum int
}

func initLoginFunctions(m *macaron.Macaron){
	m.Get("/login", login)
	m.Get("/install/admin", initAdmin)
	m.Get("/addUser", initUser)
	m.Post("/install/admin", binding.Bind(userAdd{}), addGlobalAdmin)
}

func login(ctx *macaron.Context){
	ctx.HTML(200, "login")
}

func initUser(ctx *macaron.Context, db *gorm.DB){
	ctx.HTML(200, "addUser")
}

func initAdmin(ctx *macaron.Context, tokens oauth2.Tokens, db *gorm.DB){
	if(!isLogedIn(tokens)){
		ctx.Redirect("/login")
	}
	var num int
	db.Table("users").Count(&num)
	if num != 0{
		ctx.Redirect("/")
	}
	ctx.Data["admin"] = true
	ctx.HTML(200, "addUser")
}

func isFirstLogin(sess session.Store, db *gorm.DB) bool {
	var tmpUsr database.User
	suser, ok := sess.Get("user").(map[string]interface {})
	if ok {
		db.Where("auth_id = ?", suser["user_id"].(string)).Find(&tmpUsr)
		if tmpUsr.ID == 0{
			return true
		}
		sess.Set("users", tmpUsr)
	}
	return false
}

func addGlobalAdmin(tmpUser userAdd, ctx *macaron.Context, sess session.Store, tokens oauth2.Tokens, db *gorm.DB){
	if(!isLogedIn(tokens)){
		ctx.Redirect("/login")
	}
	var num int
	db.Table("users").Count(&num)
	if num != 0{
		ctx.Redirect("/")
	}
	if newUserValidation(&tmpUser){
		var newUser database.User
		suser, _ := sess.Get("user").(map[string]interface{})
		newUser.FirstName = tmpUser.Fname
		newUser.LastName = tmpUser.Lname
		newUser.ANumber = tmpUser.Anum
		newUser.AuthId = suser["user_id"].(string)
		newUser.Email = suser["email"].(string)
		newUser.IsGlobalAdmin = true
		db.Create(&newUser)
		ctx.Redirect("/")
	}
	ctx.Data["admin"] = true
	ctx.Data["error"] = true
	ctx.HTML(200, "addUser")
}

func newUserValidation(tmpUsr *userAdd)bool{
	fMatch, ferr := regexp.MatchString("^[\\w']{2,}$", tmpUsr.Fname)
	lMatch, lerr := regexp.MatchString("^[\\w']{2,}$", tmpUsr.Lname)
	aMatch, aerr := regexp.MatchString("^[\\d]{6,}$", strconv.Itoa(tmpUsr.Anum))
	if fMatch && lMatch && aMatch && ferr == nil && lerr == nil && aerr == nil{
		return true
	}
	return false
}