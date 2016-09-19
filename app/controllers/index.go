package controllers

import (
	"fmt"

	"github.com/kataras/iris"
	"gopkg.in/redis.v4"
)

type IndexController struct {
	redis *redis.Client
}

type Response struct {
	Result  bool   `json:"result"`
	Short   string `json:"short"`
	Message string `json:"message"`
}

func init() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}

func (c *IndexController) IndexHandler(ctx *iris.Context) {
	if err := ctx.Render("index.html", nil); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func (c *IndexController) GetShortHandler(ctx *iris.Context) {
}

func (c *IndexController) ShortURLHandler(ctx *iris.Context) {
	url := ctx.FormValue("url")
	resp := new(Response)

	if string(url) == "" {
		resp.Result = false
		resp.Message = "Please input URL first..."

		ctx.JSON(iris.StatusOK, resp)
		return
	}

	fmt.Println("Input URL is:" + string(url))
}
