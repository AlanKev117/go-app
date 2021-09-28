package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlanKev117/go-app/pkg/config"
	"github.com/AlanKev117/go-app/pkg/handlers"
	"github.com/AlanKev117/go-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	// Globally create and set app config object
	var appConfig config.AppConfig

	tc, err := render.GetTemplateCache()

	if err != nil {
		log.Fatalf("error creating template cache")
	}

	// Configure app
	appConfig.TemplateCache = tc
	appConfig.UseCache = false

	// Create handler repo
	appRepository := handlers.NewRepository(&appConfig)
	handlers.SetAppRepository(appRepository)
	render.SetAppConfig(&appConfig)

	// Implement handlers
	http.HandleFunc("/", handlers.AppRepository.Home)
	http.HandleFunc("/about", handlers.AppRepository.About)

	// Run server
	fmt.Printf("Staring application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
