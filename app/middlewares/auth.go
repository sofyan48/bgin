package middlewares

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
	rd "github.com/meongbego/bgin/app/moduls/package"
)

func ChekcIPRange(w_ip []string, addr string) bool {

	for _, ip := range w_ip {
		ip_fix := strings.ReplaceAll(ip, " ", "")
		_, ipv4Net, err := net.ParseCIDR(ip_fix)
		if err != nil {
			fmt.Println(err)
		}
		if ipv4Net.Contains(net.ParseIP(addr)) {
			return true
			break
		}
	}
	return false
}

func CheckValidIP(c *gin.Context, ip string) {
	whitelist_addr := libs.GetEnvVariabel("ACL_ADDR", os.Getenv("ACL_ADDR"))
	cidr := strings.Split(whitelist_addr, ",")

	if ip == "::1" {
		c.Next()
	} else {
		check := ChekcIPRange(cidr, ip)
		if check != true {
			helper.ResponseMsg(c, 401, "Your not Authorize")
			c.Abort()
		} else {
			c.Next()
		}
	}
}

func AuthACL() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		CheckValidIP(c, ip)
	}
}

func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Access-Token"]
		_, err := redis.String(rd.Store.Do("GET", token))
		if err == redis.ErrNil {
			helper.ResponseMsg(c, 401, "Your not Authorize")
			c.Abort()
		} else if err != nil {
			panic(fmt.Sprintf("Redis Error : %v", err))
		} else {
			c.Next()
		}
	}
}
