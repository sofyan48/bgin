package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	packages "github.com/meongbego/bgin/app/moduls/package"
)

// KafkaController Types
type KafkaController struct{}

// KafkaTest Function
func (p KafkaController) KafkaTest(c *gin.Context) {
	res, _ := packages.SendMessage(packages.Kafka, "test_topic", "ok")
	helper.ResponseData(c, 200, res)
	return
}
