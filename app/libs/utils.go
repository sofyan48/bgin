package libs

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

type RespData interface{}

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

func CachData() {

}
