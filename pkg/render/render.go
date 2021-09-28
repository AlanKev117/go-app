package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/AlanKev117/go-app/pkg/config"
	"github.com/AlanKev117/go-app/pkg/models"
)

// Functions to apply to new templates
var functions = template.FuncMap{}

// App config variable
var app *config.AppConfig

// Sets the app config variable from outside this package
func SetAppConfig(appConfig *config.AppConfig) {
	app = appConfig
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// Renders a template that matches tmpl name
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = GetTemplateCache()
	}

	if tc == nil {
		log.Fatalln("could not get template cache from app config")
	}

	template, templateExists := tc[tmpl]

	if !templateExists {
		log.Fatal("template not in templates cache")
	}

	buff := new(bytes.Buffer)

	td = AddDefaultData(td)
	err := template.Execute(buff, td)
	if err != nil {
		fmt.Println(err)
	}

	_, err = buff.WriteTo(w)

	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

// Pre-compiles all templates in templates directory
func GetTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return templateCache, err
			}
		}
		templateCache[name] = ts
	}
	return templateCache, nil
}
