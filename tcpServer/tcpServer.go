package main

import (
	"fmt"
	"net"
	// only needed below for sample processing
)

func ConnHandler(c net.Conn) {

	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n]))

		_, err = c.Write(data[:n])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go ConnHandler(conn)
	}
}
