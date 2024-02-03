package main

import (
	"net/http"
	"strconv"

	"github.com/ThisJohan/HTMX-Go/views"
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())

	count := 0

	app.GET("/", func(c echo.Context) error {
		count++
		return Render(c, views.Index(strconv.Itoa(count)))
	})
	app.Logger.Fatal(app.Start(":4000"))
}

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(http.StatusOK)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
