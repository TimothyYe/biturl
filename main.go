package main

import (
	"fmt"
	"os"
	"time"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"

	"github.com/timothyye/biturl/app/controllers"
)

func main() {
	// logger and recovery middlewares are being registered inside the .Default.
	app := iris.Default()

	//Init all the settings
	initialize(web)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app.Run(iris.Addr(":" + port)) // gracefuly shutdown is enabled by-default now.
}

func initialize(app *iris.Application) {
	app.RegisterView(iris.HTML("./app/views", ".html"))

	app.WrapRouter(cors.WrapNext(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "ACCEPT", "ORIGIN"},
		AllowCredentials: true,
		Debug:            true,
	}))

	//Init controller
	indexController := &controllers.IndexController{}
	infoController := &controllers.InfoController{}

	app.Get("/", indexController.IndexHandler)
	app.Get("/{url:string}", indexController.GetShortHandler)
	app.Get("/{url:string}/info", infoController.GetURLInfoHandler)
	app.Post("/short", indexController.ShortURLHandler)
}
