package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	nodeInfo := strings.Split("120.233.10.203:6301", ":")
	fmt.Println(nodeInfo[0])
	defaultGateway := getGatewayAddr()
	fmt.Println(defaultGateway)

	err := commandExec(fmt.Sprintf("route -n |grep %s", "120.233.10.205"))
	if err != nil {
		fmt.Println("Command failed: ", err)
		return
	}
	fmt.Println("success")
	return
}

func commandExec(command string) error {
	cmd := exec.Command("bash", "-c", command)
	fmt.Println("[Exec] " + command)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Command failed: ", err)
		return err
	}
	return nil
}

func getGatewayAddr() string {
	//ip route add 11.11.11.11 via 192.168.0.168
	//ip route del 11.11.11.11
	gateway := "192.168.0.1"
	fh, err := os.Open("govpn")
	if err != nil {
		fmt.Println(err)
		return gateway
	}
	defer fh.Close()

	rd := bufio.NewReader(fh)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}

		firstRe, _ := regexp.Compile(`\s+option\s+newGateway\s+'\d+\.\d+\.\d+\.\d+'\s+`)
		gatewayLine := firstRe.FindStringSubmatch(line)
		if len(gatewayLine) == 0 {
			continue
		}

		secondRe, _ := regexp.Compile(`\d+\.\d+\.\d+\.\d+`)
		gatewayLine = secondRe.FindStringSubmatch(strings.Join(gatewayLine, ""))
		gateway = strings.Join(gatewayLine, "")
		return gateway
	}
	return gateway
}
