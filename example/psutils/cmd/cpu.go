package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"time"
)

//GetCPUInfo Show CPU Information
func GetCPUInfo() ([][]byte, error) {
	cupData := make([][]byte, 100)

	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
		return cupData, err
	}

	for _, info := range cpuInfos {
		data, _ := json.MarshalIndent(info, "", "  ")
		cupData = append(cupData, data)
	}
	return cupData, nil
}

//GetCPUPercent Show CPU usage rate
func GetCPUPercent() ([]float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return percent, err
	}
	return percent, nil
}
