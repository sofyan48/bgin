package moduls

import (
	"log"
	"time"

	"go.etcd.io/etcd/client"
)

var Etcd client.KeysAPI

// Initetcd Function
func Initetcd() client.KeysAPI {
	cfg := client.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return client.NewKeysAPI(c)
}
