package Controller

import (
	"fmt"
	"github.com/labstack/echo/v4"

	"net/http"
	"sweaty/Database"
	"sweaty/Database/Migration"
)

func CreateSweeter(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	fmt.Println(name, email)

	Db := Database.ConnectToDB().Database
	err := Db.Create(&Migration.Sweater{Email: email, Name: name})
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
	}
	return c.String(http.StatusOK , "great")
}
