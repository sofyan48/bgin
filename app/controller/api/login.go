package controller

import (
	"github.com/garyburd/redigo/redis"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/models"
	scheme "github.com/meongbego/bgin/app/moduls/migration"

	"encoding/json"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	rd "github.com/meongbego/bgin/app/moduls/package"
)

type LoginController struct{}

type Login struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int) string {
	b := make([]byte, length)
	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (h LoginController) LoginUsers(c *gin.Context) {
	var data Login
	var logindata scheme.LoginScheme

	type ResponseData struct {
		Username string `json:"username"`
		Token    string `json:"status"`
		Expire   int    `json:expire`
	}

	c.Bind(&data)
	err := models.GetByUsername(&logindata, data.Username)

	if err != nil {
		helper.ResponseMsg(c, 404, "Login Not Success")
	} else {
		if logindata.Password == data.Password {
			token := StringWithCharset(100)
			data, _ := json.Marshal(logindata)
			redis.String(rd.Store.Do("SET", data, token))
			redis.String(rd.Store.Do("EXPIRE", token, 3600))
			var response ResponseData
			response.Token = token
			response.Username = logindata.Username
			response.Expire = 3600
			helper.ResponseSuccess(c, 200, response)
		} else {
			helper.ResponseMsg(c, 404, "Username Or Password Wrong")
		}
	}
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
