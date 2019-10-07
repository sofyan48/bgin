package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	type Respons struct {
		Memory interface{} `json:"mem"`
		Cpu    interface{} `json:"cpu"`
		Disk   interface{} `json:"disk"`
	}
	var res Respons
	res.Memory = libs.GetMemHealth()
	res.Cpu = libs.GetCPU()
	res.Disk = libs.GetDiskInfo()
	helper.ResponseData(c, 200, res)
	return
}

func (h HealthController) StatusCpu(c *gin.Context) {
	helper.ResponseData(c, 200, libs.GetCPU())
	return
}

func (h HealthController) StatusMem(c *gin.Context) {
	helper.ResponseData(c, 200, libs.GetMemHealth())
	return
}

func (h HealthController) StatusDisk(c *gin.Context) {
	helper.ResponseData(c, 200, libs.GetDiskInfo())
	return
}
