package controller

import (
	"github.com/meongbego/bgin/app/helper"

	"github.com/gin-gonic/gin"
)

type PingController struct{}

func (p PingController) Status(c *gin.Context) {
	helper.ResponseSuccess(c, 200, "I'm OK!")
	return
}
