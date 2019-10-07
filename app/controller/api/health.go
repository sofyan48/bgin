package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	libs.GetCPUHealth()
	helper.ResponseMsg(c, 200, "I'm OK!")
	return
}
