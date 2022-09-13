package Route

import (
	"github.com/gin-gonic/gin"
	"sweaty/Controller"
)

func Route()  {
	router:= gin.Default()

	// render html files
	router.LoadHTMLGlob("./Resource/*.html")

	// TODO :: if start with "api" go to api function
	//api(router/)

	//start routes ************
	router.GET("/" , Controller.HomePage)
	router.GET("/about-us" , Controller.AboutUs)

	router.POST("/submit-request" ,Controller.SubmitRequest)

	//end routes   ************

	router.Run("localhost:9000")

}