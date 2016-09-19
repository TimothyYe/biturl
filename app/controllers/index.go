package controllers

import (
	"fmt"

	"github.com/kataras/iris"
)

type IndexController struct {
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
}
