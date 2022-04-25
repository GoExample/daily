package main

import (
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"time"
)

func main() {
	dial, err := net.Dial("tcp", ":3306")
	if err != nil {
		fmt.Println("Dial failure")
		return
	}
	tcpConn := dial.(*net.TCPConn)
	tcpConn.SetKeepAlive(true)
	tcpConn.SetKeepAlivePeriod(time.Second * 30) //
	fmt.Printf("%+v\n", dial)
	fmt.Printf("%#v\n", dial)
	fmt.Println(dial.RemoteAddr())
	fmt.Println(strings.Repeat("#", 80))
	connIsClosed(tcpConn)
	err = tcpConn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(strings.Repeat("-", 80))
	connIsClosed(tcpConn)
	err = dial.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	connIsClosed(tcpConn)
	fmt.Println(strings.Repeat("#", 80))
	fmt.Println(dial.RemoteAddr())
	if reflect.ValueOf(dial).IsNil() {
		fmt.Println("Close success")
	}
	fmt.Println(reflect.TypeOf(dial))
	fmt.Println(reflect.ValueOf(dial).IsValid())
	fmt.Printf("%+v\n", dial)
	fmt.Printf("%#v\n", dial)
}

func connIsClosed(c *net.TCPConn) {
	c.SetReadDeadline(time.Now())
	var one []byte
	if _, err := c.Read(one); err == io.EOF {
		fmt.Printf("Client disconnect: %s\n", c.RemoteAddr())
		c.Close()
		c = nil
	} else {
		var zero time.Time
		c.SetReadDeadline(zero)
	}
}
