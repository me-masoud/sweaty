package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"sweaty/Config"
	"sweaty/Route"
)

func main() {

	//  config
	err := Config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Config.AppConfig.Server.Port)
	// init server
	server := echo.New()
	// routing
	Route.Routing(server)

	// middleware

	//start server

	server.Start(":" + Config.AppConfig.Server.Port)

}

// getAlbums responds with the list of all albums as JSON.
//func getAlbums(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, albums)
//}

