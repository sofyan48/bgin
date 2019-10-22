package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/meongbego/bgin/app/config"
	scheme "github.com/meongbego/bgin/app/moduls/migration"
	packages "github.com/meongbego/bgin/app/moduls/package"
	"github.com/meongbego/bgin/app/server"
)

func main() {
	// Setting Config
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)

	// Setting moduls
	// storing Connectvity to Conn global Variabel
	packages.Conn = packages.InitDB()
	// Initialize Migrating All Scheme
	scheme.MigrateScheme(packages.Conn)
	// storing Connectvity to Store globalConnVariabel
	packages.Store = packages.InitRedis()
	// storing Connectvity to Kafka global Variabel
	// packages.Kafka = packages.Initkafka()
	// Up server
	server.Init()
}
