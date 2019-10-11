package moduls

import (
	"time"

	"github.com/Shopify/sarama"
	"github.com/meongbego/bgin/app/libs"
	"github.com/sirupsen/logrus"
)

// KafkaProducer Methods
type KafkaProducer struct {
	Producer sarama.SyncProducer
}

// Respons Data
type Respons struct {
	Topic  interface{} `json:"topic"`
	Data   interface{} `json:"data"`
	Offset interface{} `json:"offset"`
}

// Kafka Config
var Kafka sarama.SyncProducer

// Initkafka Function
func Initkafka() sarama.SyncProducer {
	kafkaConfig := GetKafkaConfig()
	producers, err := sarama.NewSyncProducer([]string{"localhost:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
	}
	return producers
}

// SendMessage function to send message into kafka
func SendMessage(p sarama.SyncProducer, topic string, msg string) (Respons, error) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	var res Respons

	partition, offset, err := p.SendMessage(kafkaMsg)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return res, err
	}
	res.Data = msg
	res.Offset = offset
	res.Topic = topic
	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return res, err
}

// GetKafkaConfig Get Config From kafka
func GetKafkaConfig() *sarama.Config {
	username := libs.GetEnvVariabel("KAFKA_USERNAME", "")
	password := libs.GetEnvVariabel("KAFKA_PASSWORD", "")
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
