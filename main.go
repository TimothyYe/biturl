package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/TimothyYe/biturl/app/controllers"
)

func main() {
	// logger and recovery middlewares are being registered inside the .Default.
	app := gin.Default()

	//Init all the settings
	initialize(app)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app.Run(":" + port) // gracefuly shutdown is enabled by-default now.
}

func initialize(app *gin.Engine) {
	app.LoadHTMLGlob("./app/views/*")

	//Init controller
	indexController := &controllers.IndexController{}
	infoController := &controllers.InfoController{}

	app.GET("/", indexController.IndexHandler)
	app.GET("/:url", indexController.GetShortHandler)
	app.GET("/:url/info", infoController.GetURLInfoHandler)
	app.POST("/short", indexController.ShortURLHandler)
}
