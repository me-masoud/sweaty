package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func HomePage(c *gin.Context){
	//http.FileServer(http.Dir("../Resource/home"))
	// c.IndentedJSON(200 ,"data")
	//c.HTML(http.StatusOK ,"home.html" ,gin.H{
	//	"content": "This is an home page...",
	//} )
	c.HTML(http.StatusOK ,"home.html" ,nil)
}

func AboutUs(c *gin.Context)  {
	c.HTML(http.StatusOK ,"about-us.html" ,nil)
}

func SubmitRequest(c *gin.Context)  {
	email := c.PostForm("email")

	DataBase()

	c.HTML(http.StatusOK ,"success.html" ,gin.H{
		"email" : email,
	})
}
