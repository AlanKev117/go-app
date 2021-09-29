package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AlanKev117/go-app/pkg/config"
	"github.com/AlanKev117/go-app/pkg/handlers"
	"github.com/AlanKev117/go-app/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	// Get template cache for app config
	tc, err := render.GetTemplateCache()
	if err != nil {
		log.Fatalf("error creating template cache")
	}

	// Create a session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	// Configure app
	appConfig.InProduction = false
	appConfig.TemplateCache = tc
	appConfig.UseCache = false
	appConfig.Session = session

	// Create handler repo
	handlerRepo := handlers.NewHandlerRepository(&appConfig)
	handlers.SetHandlerRepository(handlerRepo)
	render.SetAppConfig(&appConfig)

	// Create server with handlers
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	// Run server
	fmt.Printf("Staring application on port %s\n", portNumber[1:])
	err = server.ListenAndServe()
	log.Fatalln(err)
}
