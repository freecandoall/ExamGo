package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings" // only needed below for sample processing
)

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {

		// will listen for message to process ending in newline (\n)
		reader := bufio.NewReader(conn)

		line, err := reader.ReadString('\n')
		if err != nil {

			fmt.Fprintln(os.Stderr, err)

			if err == io.EOF {
				break
			}
		}

		if len(line) == 0 {
			fmt.Println("Received empty...")
			continue
		}

		// output message received
		fmt.Print("Message Received:", string(line))

		// sample process for string received
		newmessage := strings.ToUpper(line)

		// send new string back to client
		conn.Write([]byte(newmessage + "\n"))
	}
}
