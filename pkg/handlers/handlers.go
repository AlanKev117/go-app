package handlers

import (
	"net/http"

	"github.com/AlanKev117/go-app/pkg/config"
	"github.com/AlanKev117/go-app/pkg/models"
	"github.com/AlanKev117/go-app/pkg/render"
)

var Repository *HandlerRepository

type HandlerRepository struct {
	AppConfig *config.AppConfig
}

func NewHandlerRepository(appConfig *config.AppConfig) *HandlerRepository {
	return &HandlerRepository{
		AppConfig: appConfig,
	}
}

func SetHandlerRepository(repository *HandlerRepository) {
	Repository = repository
}

// Home is the handler for the home page
func (hr *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
	remoteAddress := r.RemoteAddr
	hr.AppConfig.Session.Put(r.Context(), "remote_addr", remoteAddress)
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (hr *HandlerRepository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	remoteAddress := hr.AppConfig.Session.GetString(r.Context(), "remote_addr")
	stringMap["remote_address"] = remoteAddress
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
