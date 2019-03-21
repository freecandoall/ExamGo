package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

// 인풋 결과에 대한 에러 리턴
func WaitInput(conn net.Conn, wg *sync.WaitGroup, stopChan chan bool) {

	defer wg.Done()

	for {
		select {
		default:
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Text to send: ")

			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("WaitInput() Error ReadString. " + fmt.Sprint(err))
				stopChan <- true
				return
			}

			fmt.Fprintf(conn, text+"\n")
		case <-stopChan:
			fmt.Printf("WaitInput() Stopped chan")
			return
		}
	}
}

func WaitReceive(conn net.Conn, wg *sync.WaitGroup, stopChan chan bool) {

	defer wg.Done()

	for {
		select {
		default:
			message, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				fmt.Printf("WaitReceive() Error NewReader. " + fmt.Sprint(err))
				stopChan <- true
				return
			}

			fmt.Print("Message from server: " + message)
		case <-stopChan:
			fmt.Printf("WaitReceive() Stopped chan")
			return
		}

	}
}

func main() {

	var wg sync.WaitGroup

	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	wg.Add(2)
	var stopChan = make(chan bool)

	go WaitInput(conn, &wg, stopChan)
	go WaitReceive(conn, &wg, stopChan)

	wg.Wait()

	fmt.Println("Finished...")
}
