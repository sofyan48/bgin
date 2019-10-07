package libs

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type RespData interface{}

func GetEnvVariabel(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func GetMemHealth() RespData {
	RespData, _ := mem.VirtualMemory()
	return RespData
}

func GetCPU() RespData {
	RespData, _ := cpu.Info()
	return RespData
}

func GetDiskInfo() RespData {
	RespData, _ := disk.Usage("/")
	fmt.Println(RespData)
	return RespData
}
