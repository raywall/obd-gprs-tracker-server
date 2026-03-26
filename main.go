package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	port := "5001"
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Server started! Listening on port " + port)
	fmt.Println("Point your device to your Public IP at this port.")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("New connection from: %s\n", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	for {
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Device disconnected: %s\n", conn.RemoteAddr().String())
			return
		}

		data := buf[:size]
		// Print raw Hexadecimal (useful for debugging the protocol)
		fmt.Printf("[%s] Raw: %x\n", conn.RemoteAddr().String(), data)

		// If it's a standard ASCII-based tracker (like TK103), print string
		fmt.Printf("[%s] String: %s\n", conn.RemoteAddr().String(), string(data))
	}
}
