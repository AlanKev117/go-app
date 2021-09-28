package config

import "html/template"

// This struct holds config data such as template cache
type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache      bool
}
