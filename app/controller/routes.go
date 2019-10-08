package controller

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meongbego/bgin/app/controller/api"
	"github.com/meongbego/bgin/app/middlewares"
)

func RoutesController(r *gin.Engine) {
	// Get Controller
	login := new(controller.LoginController)
	ping := new(controller.PingController)
	health := new(controller.HealthController)

	// Create Routes No Auth Declare Here
	routes := r.Group("api")
	{
		routes.POST("/login", login.LoginUsers)
	}
	// Create Routes With Auth Declare Here
	r.Use(middlewares.AuthACL())
	// r.Use(middlewares.AuthToken())
	api := r.Group("api")
	{
		api.GET("/ping", ping.Ping)
		api.GET("/kafka", ping.KafkaTest)
		api.GET("/health", health.Status)
		api.GET("/health/cpu", health.StatusCpu)
		api.GET("/health/mem", health.StatusMem)
		api.GET("/health/disk", health.StatusDisk)
		api.GET("/login/list", login.ListLogin)
	}
}
