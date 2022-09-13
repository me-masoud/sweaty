package Route

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Routing(route *echo.Echo) error {
	route.GET("/home" , func(c echo.Context) error {
		return c.String(http.StatusOK , "this is home")
	})

	userRoute := route.Group("users")
	userRoute.GET("/:id" , func(c echo.Context) error {
		return c.String(http.StatusOK , "your id is :")
	})

	return nil
}