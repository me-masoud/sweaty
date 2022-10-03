package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"net/http"
)

func HomePage(c echo.Context) error {

	fs := http.FileServer(http.Dir("./Resource"))
	http.Handle("/", fs)
	return c.Inline("./Resource", "name")

	//http.FileServer(http.Dir("../Resource/home"))
	// c.IndentedJSON(200 ,"data")
	//c.HTML(http.StatusOK ,"home.gohtml" ,gin.H{
	//	"content": "This is an home page...",
	//} )

	//c.HTML(http.StatusOK, "home.gohtml", nil)

	//tpl, err := template.ParseFiles("./Resource/home.gohtml")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//type User struct {
	//	Name   string
	//	Coupon string
	//	Amount int64
	//}
	//user := User{
	//	Name:   "Rick",
	//	Coupon: "IAMAWESOMEGOPHER",
	//	Amount: 5000,
	//}
	//err = tpl.Execute(os.Stdout, nil)
	//if err != nil {
	//	panic(err)
	//}
	//return err
}

func AboutUs(c *gin.Context) {
	c.HTML(http.StatusOK, "about-us.html", nil)
}

func SubmitRequest(c *gin.Context) {
	email := c.PostForm("email")

	c.HTML(http.StatusOK, "success.html", gin.H{
		"email": email,
	})
}
