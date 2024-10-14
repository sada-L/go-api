package controller

import (
	"github.com/gin-gonic/gin"
	"go-server/bootstrap"
	"net/http"
)

type HomeController struct {
	Env *bootstrap.Env
}

func (hc HomeController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "GO API",
		"data":  hc.Env,
	})
	return
}
