package main

import (
	"net/http"

	"github.com/AlanKev117/go-app/pkg/config"
	"github.com/AlanKev117/go-app/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(appConfig *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)
	mux.Get("/", handlers.Repository.Home)
	mux.Get("/about", handlers.Repository.About)
	return mux
}
