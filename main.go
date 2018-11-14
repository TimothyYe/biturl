package main

import (
	"os"

	"github.com/kataras/iris"

	"github.com/timothyye/biturl/app/controllers"
)

func main() {
	// logger and recovery middlewares are being registered inside the .Default.
	app := iris.Default()

	//Init all the settings
	initialize(app)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app.Run(iris.Addr(":" + port)) // gracefuly shutdown is enabled by-default now.
}

func initialize(app *iris.Application) {
	app.RegisterView(iris.HTML("./app/views", ".html"))

	//Init controller
	indexController := &controllers.IndexController{}
	infoController := &controllers.InfoController{}

	app.Get("/", indexController.IndexHandler)
	app.Get("/{url:string}", indexController.GetShortHandler)
	app.Get("/{url:string}/info", infoController.GetURLInfoHandler)
	app.Post("/short", indexController.ShortURLHandler)
}
