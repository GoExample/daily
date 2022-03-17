package main

import (
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var b = make([]byte, 1024)

	n, err := conn.Read(b)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	log.Println(string(b[:n]))
	os.Exit(0)
}
