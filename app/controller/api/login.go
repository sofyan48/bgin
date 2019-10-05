package controller

import (
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/models"
	scheme "github.com/meongbego/bgin/app/moduls/migration"

	"github.com/gin-gonic/gin"
)

type LoginController struct{}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h LoginController) LoginUsers(c *gin.Context) {
	var data Login
	c.Bind(&data)
	helper.ResponseSuccess(c, 200, data)
	return
}

func (h LoginController) ListLogin(c *gin.Context) {
	var logindata []scheme.LoginScheme
	err := models.GetAllLogin(&logindata)
	if err != nil {
		helper.ResponseMsg(c, 404, logindata)
	} else {
		helper.ResponseSuccess(c, 200, logindata)
	}
	return
}
