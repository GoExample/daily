package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	operate string
)

func usage() {
	fmt.Println(os.Args[0])
	flag.PrintDefaults()
}

func example() {
	fmt.Println("Operating network chain, please enter the correct parameters. Example:")
	fmt.Println(os.Args[0] + " -o status")
	fmt.Println("    example: Show usage")
	fmt.Println("    start: Start a network chain")
	fmt.Println("    stop: Stop a network chain")
	fmt.Println("    status: View network chain status")
}
func init() {
	flag.StringVar(&operate, "o", "example", "Operation type: /example /start /stop /status")
}

func main() {
	flag.Parse()
	flag.Usage = usage

	switch operate {
	case "start":
		pid := statusChain("./server", "0.0.0.0", 9000, false)
		if pid == 0 {
			startChain("./server", "0.0.0.0", 9000)
		}
		statusChain("./server", "0.0.0.0", 9000, true)
	case "stop":
		pid := statusChain("./server", "0.0.0.0", 9000, false)
		if pid > 0 {
			stopChain("./server", "0.0.0.0", 9000)
		}
		statusChain("./server", "0.0.0.0", 9000, true)
	case "status":
		statusChain("./server", "0.0.0.0", 9000, true)
	default:
		example()
	}
}

func startChain(server, host string, port int) {
	pendingCmd := fmt.Sprintf("%s %s %d &", server, host, port)
	cmd := exec.Command("/bin/bash", "-c", pendingCmd)
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Cmd is <%s>, err is <%s>\n", pendingCmd, err)
		return
	}
}

func stopChain(server, host string, port int) {
	pendingCmd := fmt.Sprintf("ps -ef|grep \"%s %s %d\" |grep -v grep |awk '{print $2}'|xargs kill -9", server, host, port)
	cmd := exec.Command("/bin/bash", "-c", pendingCmd)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Cmd is <%s>, err is <%s>\n", pendingCmd, err)
		return
	}
}

func statusChain(server, host string, port int, printFlag bool) int {
	pendingCmd := fmt.Sprintf("ps -ef|grep \"%s %s %d\" |grep -v grep |awk '{print $2}'", server, host, port)
	cmd := exec.Command("/bin/bash", "-c", pendingCmd)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Cmd is <%s>, err is <%s>\n", pendingCmd, err)
		return -1
	}

	pid := -1
	promptMsg := ""
	processID := strings.TrimSpace(string(output))
	if processID == "" {
		pid = 0
		promptMsg = "Server is stopped\n"
	} else {
		promptMsg = fmt.Sprintf("Server is running, process id is %s\n", processID)
		if printFlag {
			fmt.Printf("Server is running, process id is %s\n", processID)
		}
		pid, err = strconv.Atoi(processID)
		if err != nil {
			fmt.Printf("Server is running, but get process id failed, error msg is %s", err)
			return -1
		}
	}

	if printFlag {
		fmt.Printf(promptMsg)
	}
	return pid
}
