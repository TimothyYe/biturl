package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/graceful"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recovery"
)

func main() {
	web := iris.New()

	web.Use(logger.Default())
	web.Use(recovery.New(os.Stderr))

	web.Config().Render.Template.Directory = "./app/views"

	web.Get("/", hi)
	web.Static("/assets", "./app/assets", 1)

	//Get port from environment variables, default port number is 7000
	port := os.Getenv("PORT")

	if port == "" {
		port = "7000"
	}

	fmt.Println("Service is listening at:" + port)
	graceful.Run(":"+port, time.Duration(10)*time.Second, web)
}

func hi(ctx *iris.Context) {
	if err := ctx.Render("index.html", nil); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}
