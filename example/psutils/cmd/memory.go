package cmd

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v3/mem"
)

//GetSwapMemory Show system swap memory information
func GetSwapMemory() ([]byte, error) {
	memData := make([]byte, 100)
	memory, err := mem.SwapMemory()
	if err != nil {
		return memData, err
	}
	memData, _ = json.MarshalIndent(memory, "", "  ")
	return memData, nil
}

//GetVirtualMemory Show system virtual memory information
func GetVirtualMemory() ([]byte, error) {
	memData := make([]byte, 100)

	memory, err := mem.VirtualMemory()
	if err != nil {
		return memData, err
	}
	memData, _ = json.MarshalIndent(memory, "", "  ")
	return memData, nil
}
