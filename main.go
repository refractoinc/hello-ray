package main

import (
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	helloTemplate := template.New("hello_page")
	t := &Template{
		templates: template.Must(helloTemplate.Parse(hello_page)),
	}
	if rayName == "" {
		rayName = "App"
	}

	e := echo.New()
	e.Renderer = t
	e.GET("/", Hello)
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "hello_page", map[string]interface{}{
			"name": rayName,
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
