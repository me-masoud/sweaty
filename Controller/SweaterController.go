package Controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"sweaty/Model/Sweater"

	"net/http"
	"sweaty/Database"
)

func CreateSweeter(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	fmt.Println(name, email)

	Db := Database.ConnectToDB().Database

	err := Db.Create(&Sweater.Sweater{Email: email, Name: name})

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	return c.String(http.StatusOK, "great")
}

func AllSweater(c echo.Context) error {
	Db := Database.ConnectToDB().Database

	err := Db.First(&Sweater.Sweater{Email: "work.masoud@gmail.com"})
	if err != nil {
		c.String(http.StatusNotFound, "not fount dadash")
	}
	return c.JSON(http.StatusOK, err)
}
