package app

import (
	"fmt"
	"net"
)

func TCPListener(host string, port int) {
	address := fmt.Sprintf("%s:%d", host, port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer ln.Close()
	fmt.Printf("Listening for connections on %s\n", address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Printf("Connected by %s\n", conn.RemoteAddr().String())
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Printf("Received: %s\n", string(buf))
		_, err = conn.Write(buf)
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}
