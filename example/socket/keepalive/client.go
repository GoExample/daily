package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")

	err := conn.(*net.TCPConn).SetKeepAlive(true)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = conn.(*net.TCPConn).SetKeepAlivePeriod(30 * time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
	notify := make(chan error)

	go func() {
		buf := make([]byte, 1024)
		buf = []byte("my is client")
		for {
			n, err := conn.Write(buf)
			if err != nil {
				notify <- err
				if io.EOF == err {
					return
				}
			}
			if n > 0 {
				fmt.Printf("unexpected data: %s\n", buf[:n])
			}
		}
	}()

	for {
		select {
		case err := <-notify:
			fmt.Println("connection dropped message", err)
			return
		case <-time.After(time.Second * 1):
			fmt.Println("timeout 1, still alive")
		}
	}
}
