package main

import (
	"cakeapp/pkg/config"
	"cakeapp/pkg/handlers"
	"cakeapp/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Printf("Starting application on port: %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
