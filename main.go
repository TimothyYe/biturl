package main

import (
	"fmt"
	"os"
	"time"

	"github.com/iris-contrib/graceful"
	"github.com/iris-contrib/middleware/cors"
	"github.com/iris-contrib/middleware/logger"
	"github.com/iris-contrib/middleware/recovery"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
	"github.com/timothyye/biturl/app/controllers"
)

func main() {
	web := iris.New()

	//Init all the settings
	initialize(web)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "7000"
	}

	fmt.Println("Service is listening at:" + port)
	graceful.Run(":"+port, time.Duration(10)*time.Second, web)
}

func initialize(web *iris.Framework) {
	web.Use(logger.New())
	web.Use(recovery.Handler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "ACCEPT", "ORIGIN"},
		AllowCredentials: true,
		Debug:            true,
	})

	web.Use(c)
	web.UseTemplate(html.New()).Directory("./app/views", ".html")

	//Init controller
	indexController := &controllers.IndexController{}

	web.Get("/", indexController.IndexHandler)
	web.Get("/:url", indexController.GetShortHandler)
	web.Post("/short", indexController.ShortURLHandler)
}
