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
	packages.Conn = packages.InitDB()
	scheme.MigrateScheme(packages.Conn)
	packages.Store = packages.InitRedis()
	// Up server
	server.Init()
}
