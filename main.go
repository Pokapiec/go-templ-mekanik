package main

import (
	"bytes"
	"context"
	"net/http"
	"templ_test/components"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static/", "assets")
	e.GET("/", func(c echo.Context) error {
		output := bytes.Buffer{}
		component := components.Register()
		components.BaseLayout(component).Render(context.Background(), &output)
		return c.HTML(http.StatusOK, output.String())
	})

	e.POST("/register", func(c echo.Context) error {
		appRole := c.FormValue("app-role")
		username := c.FormValue("username")

		output := bytes.Buffer{}
		components.Dashboard(appRole, username).Render(context.Background(), &output)
		return c.HTML(http.StatusOK, output.String())
	})

	e.Logger.Fatal(e.Start(":3030"))
}
