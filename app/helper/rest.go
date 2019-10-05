package helper

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status   int         `json:"status"`
	Messages interface{} `json:"message"`
	Data     interface{} `json:"data"`
}

func ResponseData(w *gin.Context, status int, payload interface{}) {
	var res Response

	res.Status = status
	res.Data = payload

	w.JSON(status, res)
}

func ResponseMsg(w *gin.Context, status int, message interface{}) {
	var res Response

	res.Status = status
	res.Messages = message

	w.JSON(status, res)
}
