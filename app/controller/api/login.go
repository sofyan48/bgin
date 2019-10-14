package controller

import (
	"github.com/garyburd/redigo/redis"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
	"github.com/meongbego/bgin/app/models"
	scheme "github.com/meongbego/bgin/app/moduls/migration"

	"encoding/json"

	"github.com/gin-gonic/gin"
	rd "github.com/meongbego/bgin/app/moduls/package"
)

type LoginController struct{}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func (h LoginController) LoginUsers(c *gin.Context) {
	var data Login
	var logindata scheme.LoginScheme

	type ResponseData struct {
		Username string `json:"username"`
		Token    string `json:"token"`
		Expire   int    `json:"expire"`
	}

	c.Bind(&data)
	err := models.GetByUsername(&logindata, data.Username)

	if err != nil {
		helper.ResponseMsg(c, 404, "Login Not Success")
	} else {
		if logindata.Password == data.Password {
			token := libs.StringWithCharset(100)
			data, _ := json.Marshal(logindata)
			redis.String(rd.Store.Do("SET", token, data))
			redis.String(rd.Store.Do("EXPIRE", token, 3600))
			var response ResponseData
			response.Token = token
			response.Username = logindata.Username
			response.Expire = 3600
			helper.ResponseData(c, 200, response)
		} else {
			helper.ResponseMsg(c, 404, "Username Or Password Wrong")
		}
	}
	return
}

func (h LoginController) ListLogin(c *gin.Context) {

	value, rd_err := redis.String(rd.Store.Do("GET", "loginlist"))
	if rd_err != nil {
		var logindata []scheme.LoginScheme
		err := models.GetAllLogin(&logindata)
		if err != nil {
			helper.ResponseMsg(c, 404, logindata)
		}
		data, _ := json.Marshal(logindata)
		redis.String(rd.Store.Do("SET", "loginlist", data))
		redis.String(rd.Store.Do("EXPIRE", "loginlist", 1200))
		helper.ResponseData(c, 200, logindata)
	} else {
		type RespData []scheme.LoginScheme
		data := &RespData{}
		err := json.Unmarshal([]byte(value), data)
		if err != nil {
			helper.ResponseMsg(c, 404, err)
		} else {
			helper.ResponseData(c, 200, data)
		}
	}

	return
}
