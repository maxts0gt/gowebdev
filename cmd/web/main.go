package main

import (

	"fmt"
	"log"
	"web-dev-go/pkg/config"
	"web-dev-go/pkg/handlers"
	"web-dev-go/pkg/render"

	"net/http"
)

 const portNumber = ":8080"


 // main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

http.HandleFunc("/", handlers.Repo.Home)
http.HandleFunc("/about", handlers.Repo.About)

fmt.Println(fmt.Sprintf("Starting application on part %s", portNumber))
_ = http.ListenAndServe(portNumber, nil)
}

























































