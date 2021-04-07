package router

import (
	"net/http"

	"Server.go/controller"
	"github.com/labstack/echo"
)

// "Server.go/controller"
// "github.com/labstack/echo/v4"

func Routes() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "server đây nè")
	})
	e.POST("/user", controller.CreateStudent)
	e.GET("/users", controller.GetStudent)
	e.GET("/allusers", controller.GetAllStudent)
	e.DELETE("/delete", controller.DeleteStudent)
	e.PUT("/update", controller.EditStudent)
	return e

}
