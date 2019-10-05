package server

import (
	controller "bgin/app/controller"
	"bgin/app/libs"
	"fmt"
	"os"
)

func Init() {
	r := Routes()
	controller.RoutesController(r)
	port := libs.GetEnvVariabel("APP_PORT", os.Getenv("APP_PORT"))
	host := libs.GetEnvVariabel("APP_HOST", os.Getenv("APP_HOST"))
	r.Run(fmt.Sprintf("%s:%s", host, port))
}
