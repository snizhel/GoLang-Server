package router

import (
	"net/http"

	// "github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func Index() *echo.Echo {
	e := echo.New()
	// e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "server đây nè")
	})

	g := e.Group("/student")
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "student đây này")
	})

	return e
}
