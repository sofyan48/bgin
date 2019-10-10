package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
)

type PingController struct{}

func (p PingController) Ping(c *gin.Context) {
	helper.ResponseMsg(c, 200, "Pong !")
	return
}
