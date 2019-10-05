package middlewares

import (
	"fmt"
	"os"
	"strings"
	"net"
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
)

func ChekcIPRange(w_ip []string, addr string) bool {
	
	for _,ip := range w_ip {
		ip_fix := strings.ReplaceAll(ip, " ", "")
		_, ipv4Net, err := net.ParseCIDR(ip_fix)
		if err != nil {
			fmt.Println(err)
		}
		if ipv4Net.Contains(net.ParseIP(addr)){
			return true
			break
		}
	}
	return false
}

func CheckValidIP(c *gin.Context,ip string) {
	whitelist_addr := libs.GetEnvVariabel("ACL_ADDR", os.Getenv("ACL_ADDR"))
	cidr := strings.Split(whitelist_addr,",")
	
	if (ip == "::1"){
		c.Next()
	} else{
		check := ChekcIPRange(cidr, ip)
		if (check != true){
			helper.ResponseMsg(c, 401, "Not Valid IP")
			c.Abort()
		}else{
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
