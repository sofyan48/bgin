package controller

import (
	"github.com/gin-gonic/gin"
	"bgin/app/helper"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	helper.ResponseSuccess(c, 200, "I'm OK!")
	return
}
