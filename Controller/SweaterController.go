package Controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func CreateSweeter(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	fmt.Println(name, email)
	return nil
}
