package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
	packages "github.com/meongbego/bgin/app/moduls/package"
)

// KafkaController Types
type KafkaController struct{}

// SenDFormat | Sending To Kafka
type SenDFormat struct {
	Taskid string      `json:"task_id"`
	Data   interface{} `json:"data"`
}

// KafkaTest Function
func (p KafkaController) KafkaTest(c *gin.Context) {
	taskid := libs.StringWithCharset(20)
	topic := "test_topic"
	var sd SenDFormat
	sd.Data = map[string]string{
		"name": "iank",
	}
	sd.Taskid = topic + ":" + taskid
	data, _ := json.Marshal(sd)
	res, _ := packages.SendMessage(packages.Kafka, topic, string(data))
	helper.ResponseData(c, 200, res)
	return
}
