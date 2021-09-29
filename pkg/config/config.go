package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// This struct holds config data such as template cache
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
	InProduction  bool
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
