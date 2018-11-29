package controllers

import (
	"strings"

	"github.com/kataras/iris"
	"github.com/timothyye/biturl/utils"
	"gopkg.in/redis.v5"
)

var client *redis.Client

const (
	domain = "biturl.top"
	url    = "https://biturl.top/"
)

//IndexController for URL shorten handling
type IndexController struct {
	// iris.Controller
	// MVC architectural pattern is built'n inside Iris now but
	// I'll not change your design here, it will be kept it as handler-driven.
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
func (c *IndexController) IndexHandler(ctx iris.Context) {
	ctx.View("index.html")
}

//GetShortHandler for getting shorten URL querying result
func (c *IndexController) GetShortHandler(ctx iris.Context) {
	url := ctx.Params().Get("url")
	longURL := client.Get(url).Val()

	if len(longURL) > 0 {
		if strings.HasPrefix(longURL, "http://") || strings.HasPrefix(longURL, "https://") {
			ctx.Redirect(longURL)
			return
		}

		ctx.Redirect("https://" + longURL)
		return
	}

	ctx.Redirect("/")
}

//ShortURLHandler for shorten long URL
func (c *IndexController) ShortURLHandler(ctx iris.Context) {
	url := ctx.FormValue("url")
	resp := new(Response)
	inputURL := string(url)

	if !strings.HasPrefix(inputURL, "http") {
		inputURL = "https://" + inputURL
	}

	if inputURL == "" {
		resp.Result = false
		resp.Message = "Please input URL first..."

		ctx.JSON(resp)
		return
	}

	if strings.Contains(inputURL, domain) {
		resp.Result = false
		resp.Message = "Cannot shorten it again..."

		ctx.JSON(resp)
		return
	}

	urls := utils.ShortenURL(inputURL)
	err := client.Set(urls[0], inputURL, 0).Err()
	if err != nil {
		resp.Result = false
		resp.Message = "Backend service is unavailable!"
	}

	resp.Result = true
	resp.Short = url + urls[0]

	ctx.JSON(resp)
}
