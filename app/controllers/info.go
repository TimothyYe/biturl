package controllers

import "github.com/kataras/iris"

//InfoController for URL shorten handling
type InfoController struct {
}

//GetURLInfoHandler for rendering the index page
func (c *InfoController) GetURLInfoHandler(ctx iris.Context) {
	url := ctx.Params().Get("url")
	ctx.Application().Logger().Infof("Info for URL: %s", url)
	ctx.Redirect("/")
}
