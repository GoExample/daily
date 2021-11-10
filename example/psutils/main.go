package main

import (
	"fmt"
	"psutils/cmd"
)

func main() {
	hostInfo, err := cmd.GetHostInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("host info: %s\n", string(hostInfo))
	outputSepLine()

	cpuInfos, err := cmd.GetCPUInfo()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("CPU information:")
	for _, info := range cpuInfos {
		fmt.Printf(string(info))
	}
	outputSepLine()

	percent, err := cmd.GetCPUPercent()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("cpu percent:%v\n", percent)
	outputSepLine()

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
	outputSepLine()

	cmd.GetDiskInfo()
	cmd.GetNetIoInformation()
	outputSepLine()

	ip, err := cmd.GetLocalIP()
	if err != nil {
		return
	}
	fmt.Printf("Local Ip:\n%s\n", ip)
	fmt.Println("Outbound Ip: ")
	cmd.GetOutboundIP()
}

func outputSepLine() {
	fmt.Println("\n================================================================================")
}
