package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	helper.ResponseMsg(c, 200, "I'm OK!")
	return
}
