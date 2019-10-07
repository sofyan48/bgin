package moduls

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

type MessageMetadata struct {
	ReceivedAt time.Time `json:"received_at"`
}

type Message struct {
	Partition int32           `json:"partition"`
	Offset    int64           `json:"offset"`
	Value     string          `json:"value"`
	Metadata  MessageMetadata `json:"metadata"`
}

type KafkaClient struct {
	producer sarama.AsyncProducer
	consumer *cluster.Consumer

	ml               sync.RWMutex
	receivedMessages []Message
}

type AppConfig struct {
	Kafka struct {
		URL           string `env:"KAFKA_URL,required"`
		TrustedCert   string `env:"KAFKA_TRUSTED_CERT,required"`
		ClientCertKey string `env:"KAFKA_CLIENT_CERT_KEY,required"`
		ClientCert    string `env:"KAFKA_CLIENT_CERT,required"`
		Prefix        string `env:"KAFKA_PREFIX"`
		Topic         string `env:"KAFKA_TOPIC,default=messages"`
		ConsumerGroup string `env:"KAFKA_CONSUMER_GROUP,default=heroku-kafka-demo-go"`
	}

	Web struct {
		Port string `env:"PORT,required"`
	}
}

func (ac *AppConfig) createKafkaConsumer(brokers []string, tc *tls.Config) *cluster.Consumer {
	config := cluster.NewConfig()

	config.Net.TLS.Config = tc
	config.Net.TLS.Enable = true
	config.Group.PartitionStrategy = cluster.StrategyRoundRobin
	config.ClientID = ac.Kafka.ConsumerGroup
	config.Consumer.Return.Errors = true

	topic := ac.topic()

	log.Printf("Consuming topic %s on brokers: %s", topic, brokers)

	err := config.Validate()
	if err != nil {
		log.Fatal(err)
	}

	consumer, err := cluster.NewConsumer(brokers, ac.group(), []string{topic}, config)
	if err != nil {
		log.Fatal(err)
	}
	return consumer
}

// Create the Kafka asynchronous producer
func (ac *AppConfig) createKafkaProducer(brokers []string, tc *tls.Config) sarama.AsyncProducer {
	config := sarama.NewConfig()

	config.Net.TLS.Config = tc
	config.Net.TLS.Enable = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll // Default is WaitForLocal
	config.ClientID = ac.Kafka.ConsumerGroup

	err := config.Validate()
	if err != nil {
		log.Fatal(err)
	}
	producer, err := sarama.NewAsyncProducer(brokers, config)
	if err != nil {
		log.Fatal(err)
	}

	return producer
}

func (ac *AppConfig) brokerAddresses() []string {
	urls := strings.Split(ac.Kafka.URL, ",")
	addrs := make([]string, len(urls))
	for i, v := range urls {
		u, err := url.Parse(v)
		if err != nil {
			log.Fatal(err)
		}
		addrs[i] = u.Host
	}
	return addrs
}

// Prepends prefix to topic if provided
func (ac *AppConfig) topic() string {
	topic := ac.Kafka.Topic

	if ac.Kafka.Prefix != "" {
		topic = strings.Join([]string{ac.Kafka.Prefix, topic}, "")
	}

	return topic
}

// Prepend prefix to consumer group if provided
func (ac *AppConfig) group() string {
	group := ac.Kafka.ConsumerGroup

	if ac.Kafka.Prefix != "" {
		group = strings.Join([]string{ac.Kafka.Prefix, group}, "")
	}

	return group
}

func newKafkaClient(config *AppConfig) *KafkaClient {
	tlsConfig := config.createTLSConfig()
	brokerAddrs := config.brokerAddresses()

	// verify broker certs
	for _, b := range brokerAddrs {
		ok, err := verifyServerCert(tlsConfig, config.Kafka.TrustedCert, b)
		if err != nil {
			log.Fatal("Get Server Cert Error: ", err)
		}

		if !ok {
			log.Fatalf("Broker %s has invalid certificate!", b)
		}
	}
	log.Println("All broker server certificates are valid!")

	return &KafkaClient{
		consumer: config.createKafkaConsumer(brokerAddrs, tlsConfig),
		producer: config.createKafkaProducer(brokerAddrs, tlsConfig),
	}
}

func verifyServerCert(tc *tls.Config, caCert string, url string) (bool, error) {
	// Create connection to server
	conn, err := tls.Dial("tcp", url, tc)
	if err != nil {
		return false, err
	}

	// Pull servers cert
	serverCert := conn.ConnectionState().PeerCertificates[0]

	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(caCert))
	if !ok {
		return false, errors.New("Unable to parse Trusted Cert")
	}

	// Verify Server Cert
	opts := x509.VerifyOptions{Roots: roots}
	if _, err := serverCert.Verify(opts); err != nil {
		log.Println("Unable to verify Server Cert")
		return false, err
	}

	return true, nil
}

// Create the TLS context, using the key and certificates provided.
func (ac *AppConfig) createTLSConfig() *tls.Config {
	roots := x509.NewCertPool()
	ok := roots.AppendCertsFromPEM([]byte(ac.Kafka.TrustedCert))
	if !ok {
		log.Println("Unable to parse Root Cert:", ac.Kafka.TrustedCert)
	}

	// Setup certs for Sarama
	cert, err := tls.X509KeyPair([]byte(ac.Kafka.ClientCert), []byte(ac.Kafka.ClientCertKey))
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
		RootCAs:            roots,
	}

	tlsConfig.BuildNameToCertificate()
	return tlsConfig
}

func InitKafka() {

}
