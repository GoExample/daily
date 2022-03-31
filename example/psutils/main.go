package main

import (
	"fmt"
	"strings"

	"psutils/cmd"
)

func main() {
	hostInfo, err := cmd.GetHostInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("host info: %s\n", string(hostInfo))
	fmt.Println("\n" + strings.Repeat("=", 80))

	cpuInfos, err := cmd.GetCPUInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CPU information:")
	for _, info := range cpuInfos {
		fmt.Printf(string(info))
	}
	fmt.Println("\n" + strings.Repeat("=", 80))

	percent, err := cmd.GetCPUPercent()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("cpu percent:%v\n", percent)
	fmt.Println("\n" + strings.Repeat("=", 80))

	swapMemory, err := cmd.GetSwapMemory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("system swap memory information: %s\n", string(swapMemory))

	virtualMemory, err := cmd.GetVirtualMemory()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("system virtual memory information: %s\n", string(virtualMemory))
	fmt.Println("\n" + strings.Repeat("=", 80))

	cmd.GetDiskInfo()
	cmd.GetNetIoInformation()
	fmt.Println("\n" + strings.Repeat("=", 80))

	ip, err := cmd.GetLocalIP()
	if err != nil {
		return
	}
	fmt.Printf("Local Ip:\n%s\n", ip)
	fmt.Println("Outbound Ip: ")
	cmd.GetOutboundIP()
}
