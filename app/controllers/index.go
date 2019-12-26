package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/TimothyYe/biturl/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v5"
)

var client *redis.Client

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
	domain      = "biturl.top"
	siteURL     = "https://biturl.top/"
	visitKey    = `visit/%s`
)

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
func (c *IndexController) IndexHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

//GetShortHandler for getting shorten URL querying result
func (c *IndexController) GetShortHandler(ctx *gin.Context) {
	url := ctx.Param("url")
	url = utils.FormatURL(&url)
	longURL := client.Get(url).Val()

	if len(longURL) > 0 {
		if strings.HasPrefix(longURL, httpPrefix) || strings.HasPrefix(longURL, httpsPrefix) {
			ctx.Redirect(http.StatusTemporaryRedirect, longURL)
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, httpsPrefix+longURL)
		return
	}

	// update visit count
	client.Incr(fmt.Sprintf(visitKey, url))

	// redirect user to the target URL
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}

//ShortURLHandler for shorten long URL
func (c *IndexController) ShortURLHandler(ctx *gin.Context) {
	url := ctx.PostForm("url")
	resp := new(Response)
	inputURL := string(url)
	url = utils.FormatURL(&url)

	if !strings.HasPrefix(inputURL, "http") {
		inputURL = httpsPrefix + inputURL
	}

	if inputURL == "" {
		resp.Result = false
		resp.Message = "Please input URL first..."

		ctx.JSON(http.StatusOK, resp)
		return
	}

	if strings.Contains(inputURL, domain) {
		resp.Result = false
		resp.Message = "Cannot shorten it again..."

		ctx.JSON(http.StatusOK, resp)
		return
	}

	urls := utils.ShortenURL(inputURL)

	// save short URL in redis
	err := client.Set(urls[0], inputURL, 0).Err()
	if err != nil {
		resp.Result = false
		resp.Message = "Backend service is unavailable!"
	}

	// set default visit info
	err = client.Set(fmt.Sprintf(visitKey, urls[0]), 0, 0).Err()
	if err != nil {
		resp.Result = false
		resp.Message = "Backend service is unavailable!"
	}

	resp.Result = true
	resp.Short = siteURL + urls[0]

	ctx.JSON(http.StatusOK, resp)
}
