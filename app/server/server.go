package server

import (
	"fmt"
	"os"

	controller "github.com/meongbego/bgin/app/controller"
	"github.com/meongbego/bgin/app/libs"
)

// Init Create Server
func Init() {
	r := Routes()
	// Export routes From Controller
	controller.RoutesController(r)
	// Add Env PORT and HOST
	port := libs.GetEnvVariabel("APP_PORT", os.Getenv("APP_PORT"))
	host := libs.GetEnvVariabel("APP_HOST", os.Getenv("APP_HOST"))
	// Create Servers
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
