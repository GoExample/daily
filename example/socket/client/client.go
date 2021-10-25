package main

import (
    "log"
    "net"
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
        log.Fatal(err)
    }

    log.Println(string(b[:n]))
}
