package main

import (
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

	// init server
	server := echo.New()

	// routing
	Route.Routing(server)

	// middleware
	//start server
	server.Start(
		Config.AppConfig.Server.Domain + ":" + Config.AppConfig.Server.Port)

}

// getAlbums responds with the list of all albums as JSON.
//func getAlbums(c *gin.Context) {
//	c.IndentedJSON(http.StatusOK, albums)
//}
