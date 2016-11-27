package controllers

import (
	"fmt"
	"strings"

	"github.com/kataras/iris"
	"github.com/timothyye/biturl/utils"
	"gopkg.in/redis.v5"
)

var client *redis.Client

//IndexController for URL shorten handling
type IndexController struct {
}

//Response struct for http response
type Response struct {
	Result  bool   `json:"result"`
	Short   string `json:"short"`
	Message string `json:"message"`
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//IndexHandler for rendering the index page
func (c *IndexController) IndexHandler(ctx *iris.Context) {
	if err := ctx.Render("index.html", nil); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

//GetShortHandler for getting shorten URL querying result
func (c *IndexController) GetShortHandler(ctx *iris.Context) {
	url := ctx.Param("url")
	//fmt.Println("original url is:", client.Get(url).Val())
	ctx.Redirect(client.Get(url).Val())
}

//ShortURLHandler for shorten long URL
func (c *IndexController) ShortURLHandler(ctx *iris.Context) {
	url := ctx.FormValue("url")
	resp := new(Response)
	inputURL := string(url)

	if inputURL == "" {
		resp.Result = false
		resp.Message = "Please input URL first..."

		ctx.JSON(iris.StatusOK, resp)
		return
	}

	if strings.Contains(inputURL, "biturl.top") {
		resp.Result = false
		resp.Message = "Cannot shorten it again..."

		ctx.JSON(iris.StatusOK, resp)
		return
	}

	urls := utils.ShortenURL(inputURL)
	err := client.Set(urls[0], inputURL, 0).Err()
	if err != nil {
		resp.Result = false
		resp.Message = "Backend service is unavailable!"
	}

	resp.Result = true
	resp.Short = "http://biturl.top/" + urls[0]

	ctx.JSON(iris.StatusOK, resp)
}
