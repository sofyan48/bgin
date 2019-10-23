package moduls

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

// ElConn | Global Var to conn
var ElConn *elastic.Client

// InitElastic Init To Connect
func InitElastic() *elastic.Client {
	// elhost := libs.GetEnvVariabel("ELASTIC_HOST_PORT", "localhost:9200")
	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		// Handle error
		logrus.Errorf("Error :%v", err)
	}
	return client
}
