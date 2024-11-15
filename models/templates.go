package models

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
)

// Templates struct copied by Primagen Setup Tutorial
// Video: https://www.youtube.com/watch?v=x7v6SNIgJpE&t=1256s

type Templates struct {
	templates *template.Template
}

// Render template to the response writer.
func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// NewTemplates initializes and parses the templates from the provided path.
func NewTemplates(path string) *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob(path)),
	}
}
