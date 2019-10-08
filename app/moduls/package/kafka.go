package moduls

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

// KafkaProducer hold kafka producer session
type KafkaProducer struct {
	Producer sarama.SyncProducer
}

func InitKafka() {
	_, err := sarama.NewSyncProducer([]string{":9092"}, GetKafkaConfig())
	fmt.Println(err)
	// if err != nil {
	// 	logrus.Errorf("Unable to create kafka producer got error %v", err)
	// 	return
	// }
	// defer func() {
	// 	if err := producers.Close(); err != nil {
	// 		logrus.Errorf("Unable to stop kafka producer: %v", err)
	// 		return
	// 	}
	// }()
}

func GetKafkaConfig() *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0
	return kafkaConfig
}

// SendMessage function to send message into kafka
func SendMessage(topic, msg string) error {
	var p *KafkaProducer
	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	partition, offset, err := p.Producer.SendMessage(kafkaMsg)
	if err != nil {
		logrus.Errorf("Send message error: %v", err)
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}
