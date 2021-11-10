package cmd

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v3/host"
)

//GetHostInfo Show host information.
func GetHostInfo() ([]byte, error) {
	hostData := make([]byte, 100)
	hostInfo, err := host.Info()
	if err != nil {
		return hostData, err
	}
	hostData, err = json.MarshalIndent(hostInfo, "", "  ")
	if err != nil {
		return hostData, err
	}
	return hostData, nil
}
