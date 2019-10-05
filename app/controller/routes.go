package controller

import (
	controller "github.com/meongbego/bgin/app/controller/api"

	"github.com/gin-gonic/gin"
)

func RoutesController(r *gin.Engine) {
	api := r.Group("api")
	{
		ping := new(controller.PingController)
		health := new(controller.HealthController)
		login := new(controller.LoginController)
		// create rest api models
		api.GET("/ping", ping.Status)
		api.GET("/health", health.Status)
		api.POST("/login", login.LoginUsers)
		api.GET("/login/list", login.ListLogin)
	}
}
