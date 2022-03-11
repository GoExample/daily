package main

import (
	"context"
	"fmt"
	"golang.org/x/sys/unix"
	"net"
	"syscall"
)

func PortMultiplexing1() {
	cfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1)
			})
		},
	}
	tcp, err := cfg.Listen(context.Background(), "tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("listen failed", err)
		return
	}

	buf := make([]byte, 1024)
	for {
		conn, err := tcp.Accept()
		if err != nil {
			fmt.Println("accept failed", err)
			continue
		}
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("read failed", err)
				break
			}

			fmt.Println(string(buf[:n]))
		}
	}
}

func PortMultiplexing2(addr string) {
	cfg := net.ListenConfig{
		Control: func(network, address string, c syscall.RawConn) error {
			return c.Control(func(fd uintptr) {
				syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEADDR, 1)
			})
		},
	}
	udp, err := cfg.ListenPacket(context.Background(), "udp", addr)

	if err != nil {
		fmt.Println("listen failed", err)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, caddr, err := udp.ReadFrom(buf)
		if err != nil {
			fmt.Println("read failed", err)
			continue
		}

		fmt.Println(addr, caddr, string(buf[:n]))
	}
}
