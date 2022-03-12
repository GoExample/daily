package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	fmt.Println("建立与服务端的链接")
	conn, err := net.Dial("tcp", "127.0.0.1:7890") //创建用于通信socket
	if err != nil {
		fmt.Printf("client dial failed, err msg is %s\n", err)
		return
	}
	fmt.Println("连接成功")
	conn.Close()

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Second)
	fmt.Println(rand.Intn(rand.Intn(100)))
}
