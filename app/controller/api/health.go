package controller

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/meongbego/bgin/app/helper"
	"github.com/meongbego/bgin/app/libs"
	rd "github.com/meongbego/bgin/app/moduls/package"
)

type HealthController struct{}

// Status Get health Status All
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

	value, rderr := rd.GetRowsCached("healthstats")
	if rderr != nil {
		data, _ := json.Marshal(res)
		_, cerr := rd.RowsCached("healthstats", data, 120)
		if cerr != nil {
			fmt.Println("Not Cached")
		}
		helper.ResponseData(c, 200, res)
	} else {
		var data Respons
		err := json.Unmarshal([]byte(value), &data)
		if err != nil {
			helper.ResponseMsg(c, 404, err)
		} else {
			helper.ResponseData(c, 200, data)
		}
	}
	return
	helper.ResponseData(c, 200, res)
	return
}

// StatusCpu Get CPU Stats in spesific
func (h HealthController) StatusCpu(c *gin.Context) {
	value, rderr := rd.GetRowsCached("cpustats")
	if rderr != nil {
		data, _ := json.Marshal(libs.GetCPU())
		_, cerr := rd.RowsCached("cpustats", data, 120)
		if cerr != nil {
			fmt.Println("Not Cached")
		}
		helper.ResponseData(c, 200, libs.GetCPU())
	} else {
		var data []libs.RespData
		err := json.Unmarshal([]byte(value), &data)
		if err != nil {
			helper.ResponseMsg(c, 404, err)
		} else {
			helper.ResponseData(c, 200, data)
		}
	}
	return
}

// StatusCpu Get Mem Status in spesific
func (h HealthController) StatusMem(c *gin.Context) {
	// helper.ResponseData(c, 200, libs.GetMemHealth())
	value, rderr := rd.GetRowsCached("memstats")
	if rderr != nil {
		data, _ := json.Marshal(libs.GetMemHealth())
		_, cerr := rd.RowsCached("memstats", data, 120)
		if cerr != nil {
			fmt.Println("Not Cached")
		}
		helper.ResponseData(c, 200, libs.GetMemHealth())
	} else {
		var data []libs.RespData
		err := json.Unmarshal([]byte(value), &data)
		if err != nil {
			helper.ResponseMsg(c, 404, err)
		} else {
			helper.ResponseData(c, 200, data)
		}
	}
	return
}

// StatusCpu Get Disk Stats in spesific
func (h HealthController) StatusDisk(c *gin.Context) {
	// helper.ResponseData(c, 200, libs.GetDiskInfo())
	value, rderr := rd.GetRowsCached("diskstats")
	if rderr != nil {
		data, _ := json.Marshal(libs.GetDiskInfo())
		_, cerr := rd.RowsCached("diskstats", data, 120)
		if cerr != nil {
			fmt.Println("Not Cached")
		}
		helper.ResponseData(c, 200, libs.GetDiskInfo())
	} else {
		var data []libs.RespData
		err := json.Unmarshal([]byte(value), &data)
		if err != nil {
			helper.ResponseMsg(c, 404, err)
		} else {
			helper.ResponseData(c, 200, data)
		}
	}
	return
}
