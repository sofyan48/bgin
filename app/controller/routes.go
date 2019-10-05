package controller

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meongbego/bgin/app/controller/api"
	"github.com/meongbego/bgin/app/middlewares"
)

func RoutesController(r *gin.Engine) {
	// Create Routes No Auth
	routes := r.Group("api")
	{
		login := new(controller.LoginController)
		routes.POST("/login", login.LoginUsers)
	}
	// Create Routes With Auth
	r.Use(middlewares.AuthACL())
	r.Use(middlewares.AuthToken())
	api := r.Group("api")
	{
		ping := new(controller.PingController)
		health := new(controller.HealthController)
		login := new(controller.LoginController)

		api.GET("/ping", ping.Status)
		api.GET("/health", health.Status)
		api.GET("/login/list", login.ListLogin)
	}
}
