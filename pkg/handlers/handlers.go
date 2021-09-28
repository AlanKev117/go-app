package handlers

import (
	"net/http"

	"github.com/AlanKev117/go-app/pkg/config"
	"github.com/AlanKev117/go-app/pkg/models"
	"github.com/AlanKev117/go-app/pkg/render"
)

var AppRepository *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

func SetAppRepository(repository *Repository) {
	AppRepository = repository
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "hello from the template data"
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
