package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
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
	var sd SenDFormat
	sd.Data = map[string]string{
		"name":      "iank",
		"ticket_id": "1112222221111",
	}
	sd.Taskid = "task_123"
	data, _ := json.Marshal(sd)
	res, _ := packages.SendMessage(packages.Kafka, "test_topic", string(data))
	fmt.Println(res)
	helper.ResponseData(c, 200, res)
	return
}
