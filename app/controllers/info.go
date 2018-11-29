package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//InfoController for URL shorten handling
type InfoController struct {
}

//GetURLInfoHandler for rendering the index page
func (c *InfoController) GetURLInfoHandler(ctx *gin.Context) {
	//url := ctx.Param("url")
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
