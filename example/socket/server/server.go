package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func main() {
	stopChain := make(chan os.Signal)
	signal.Notify(stopChain, os.Interrupt)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		<-stopChain
		log.Println("Get stop command. Now stopping ...")
		if err = listen.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	log.Println("Start listen :8080: ... ")
	for {
		conn, err := listen.Accept()
		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}
			log.Println(err)
			continue
		}
		log.Println("Accept ", conn.RemoteAddr())
		wg.Add(1)
		go handler(conn)
	}
	wg.Wait()
}

func handler(conn net.Conn) {
	defer wg.Done()
	defer conn.Close()

	time.Sleep(time.Second * 3)
	conn.Write([]byte("Hello"))
	log.Println("Send Hello")
}
