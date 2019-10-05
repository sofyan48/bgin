package test

import (
	"fmt"
	"go_boilerplate/app/controller"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.TestMode)

	api := router.Group("api")
	{
		ping := new(controller.PingController)
		api.GET("/ping", ping.Status)
	}

	return router
}

func main() {
	r := SetupRouter()
	r.Run()
}

func TestPingApi(t *testing.T) {
	testRouter := SetupRouter()
	url := "/api/ping"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()

	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, 200)
}
