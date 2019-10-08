package controller

import (
	"github.com/meongbego/bgin/app/helper"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (p PingController) Ping(c *gin.Context) {
	helper.ResponseMsg(c, 200, "Pong !")
	return
}
