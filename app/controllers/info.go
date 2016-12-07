package controllers

import (
	"fmt"

	"github.com/kataras/iris"
)

//InfoController for URL shorten handling
type InfoController struct {
}

//GetURLInfoHandler for rendering the index page
func (c *InfoController) GetURLInfoHandler(ctx *iris.Context) {
	url := ctx.Param("url")
	fmt.Println("Info for URL:", url)
	ctx.Redirect("/")
}
