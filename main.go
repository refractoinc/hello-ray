package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

var (
	rayName = os.Getenv("RAY_NAME")
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"name": rayName,
	})
}

func main() {
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	if rayName == "" {
		rayName = "App"
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", Hello)

	e.Logger.Fatal(e.Start(":8080"))
}
