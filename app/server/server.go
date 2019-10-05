package server

import (
	"fmt"
	"os"

	controller "github.com/meongbego/bgin/app/controller"
	"github.com/meongbego/bgin/app/libs"
)

func Init() {
	r := Routes()
	controller.RoutesController(r)
	port := libs.GetEnvVariabel("APP_PORT", os.Getenv("APP_PORT"))
	host := libs.GetEnvVariabel("APP_HOST", os.Getenv("APP_HOST"))
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
