package main

import (
	"fmt"
	"os/exec"
)

func main() {
	host := "0.0.0.0"
	basePort := 9000
	for i := 0; i < 1000; i++ {
		pendingCmd := fmt.Sprintf("./server %s %d &", host, basePort+i)
		fmt.Println(pendingCmd)
		cmd := exec.Command("/bin/bash", "-c", pendingCmd)
		err := cmd.Start()
		if err != nil {
			fmt.Printf("Cmd is <%s>, err is <%s>\n", pendingCmd, err)
			return
		}
	}
	/*
		count := 0
		for {
			fmt.Println("Start to sleep for 5 seconds")
			time.Sleep(time.Second * 5)
			count += 1
			fmt.Printf("Sleeped for a total of %d seconds\n", count*5)
		}
	*/
}
