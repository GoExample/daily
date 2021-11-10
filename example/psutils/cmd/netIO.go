package cmd

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
)

//GetNetIoInformation Show network information
func GetNetIoInformation() {
	info, _ := net.IOCounters(true)
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}
