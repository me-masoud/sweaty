package Route

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sweaty/Controller"
	"sweaty/Database/Migration"
)

func Routing(route *echo.Echo) error {

	route.GET("/home", func(c echo.Context) error {
		return c.String(http.StatusOK, "this is home")
	})

	/**
	Start Emails Group
	*/
	emailGroupRoute := route.Group("/email")
	emailGroupRoute.POST("/new-sweater", Controller.CreateSweeter)
	/**
	End Email Group
	*/


	/**
	Start users Group
	*/
	migrationGroupRoute := route.Group("/migration")
	migrationGroupRoute.GET("/run", Migration.Migration)
	/**
	End Migration Group
	*/


	/**
	Start users Group
	*/
	userGroupRoute := route.Group("users")
	userGroupRoute.GET("/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "your id is :")
	})
	/**
	End users Group
	*/

	return nil
}
