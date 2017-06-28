package main

import (
	"gopkg.in/macaron.v1"
	"mjolnir/routes"
	"github.com/go-macaron/pongo2"
	"github.com/go-macaron/gzip"
	"mjolnir/database"
	"github.com/go-macaron/session"
	"github.com/dragonman117/oauth2"
	goauth2 "golang.org/x/oauth2"
)

/*
Need to do authentication
 */

func main(){
	//Initialize
	m := macaron.New()

	//OAuth config
	authConfig := goauth2.Config{
		ClientID:"3kXXg0OA1Xg7Hs5iZocODeVv8cKHb5Wr",
		ClientSecret: "nWsD0TCdXWaVaRwdZyZLeCaSVvhm9Tfh06t5lM-eqE393qqSniMs4_ggEt6AjUlO",
		RedirectURL: "http://127.0.0.1:3000/oauth2callback",
		Scopes:       []string{"openid", "profile"},
	}
	oauth2.PathLogin = "/oauthLogin" // kinda  an unused path with oauth...

	//Register middleware
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(gzip.Gziper(gzip.Options{CompressionLevel:4}))
	m.Use(pongo2.Pongoer())
	m.Use(session.Sessioner())
	m.Use(oauth2.AuthZero(&authConfig, "usu-dev-student.auth0.com"))
	m.Use(macaron.Static("public"))

	//init databases
	db := database.InitDatabase()
	defer database.CloseDatabase()

	//Custom Services
	m.Map(db)

	//Pull in routes
	routes.InitRoutes(m)

	//Start serving
	m.Run("127.0.0.1", 3000)
}