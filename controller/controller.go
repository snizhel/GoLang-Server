package controller

import (
	"net/http"

	"fmt"
	"strconv"

	"Server.go/models"
	"github.com/labstack/echo"
)

var Students []models.Student
var Students1 = make(map[int]models.Student)

// var StudentTest = []map[int]model.Student
func GetStudent(c echo.Context) error {
	id := c.QueryParam("id")
	fmt.Println(id)

	//fmt.Println(Studen)

	// var Students []map[string]*models.Student
	// for _, s := range id {
	// 	// if s == id {
	// 	// 	fmt.Println(s)
	// 	// }
	// 	return c.JSON(http.StatusOK, s)

	// }
	for i := 0; i < len(Students); i++ {
		if Students[i].ID == id {
			return c.JSON(http.StatusOK, Students[i])
		}
	}
	return c.String(http.StatusOK, "student not found")
}

func CreateStudent(c echo.Context) error {
	e := new(models.Student)
	if err := c.Bind(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	Students = append(Students, *e)
	return c.JSON(http.StatusCreated, "create successful")

}
func GetAllStudent(c echo.Context) error {
	return c.JSON(http.StatusOK, Students)
}
func DeleteStudent(c echo.Context) error {
	id := c.QueryParam("id")
	for i := 0; i < len(Students); i++ {
		if Students[i].ID == id {
			Students = append(Students[:i], Students[i+1:]...)
			return c.JSON(http.StatusOK, Students)
		}

	}
	return c.String(http.StatusOK, "delete successful")
}

func EditStudent(c echo.Context) error {
	e := new(models.Student)
	if err := c.Bind(e); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	id, _ := strconv.Atoi(c.Param("id"))
	Students[id].Name = e.Name
	return c.JSON(http.StatusOK, Students[id])

}
