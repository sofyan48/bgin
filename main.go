package main

import (
	"flag"
	"fmt"
	"os"

	"bgin/app/config"
	scheme "bgin/app/moduls/migration"
	packages "bgin/app/moduls/package"
	"bgin/app/server"
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
	packages.Conn = packages.InitDB()
	scheme.MigrateScheme(packages.Conn)
	packages.Store = packages.InitRedis()
	// Up server
	server.Init()
}
