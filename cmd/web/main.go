package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
	"web-dev-go/pkg/config"
	"web-dev-go/pkg/handlers"
	"web-dev-go/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
 // main is the main application function
func main() {


	app.InProduction = false
	// in production, change it to true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false //in development mode it should be false. since we are hitting localhost:8080

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)



fmt.Println(fmt.Sprintf("Starting application on part %s", portNumber))

srv := &http.Server {
	Addr: portNumber,
	Handler: routes(&app),
}


 err = srv.ListenAndServe()
 log.Fatal(err)
}























































