package controller

import (
	"github.com/Shopify/sarama"
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	producer "github.com/meongbego/bgin/app/moduls/package"
	"github.com/sirupsen/logrus"
)

// KafkaController Types
type KafkaController struct{}

// KafkaTest Function
func (p KafkaController) KafkaTest(c *gin.Context) {
	kafkaConfig := producer.GetKafkaConfig()
	producers, err := sarama.NewSyncProducer([]string{"localhost:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
	}
	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
		}
	}()
	kafka := &producer.KafkaProducer{
		Producer: producers,
	}
	kafka.SendMessage("test_topic", "test")
	helper.ResponseMsg(c, 200, "")
	return
}
