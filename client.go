package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	go listenServer(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		//fmt.Print("Enter message: ")
		message, _ := reader.ReadString('\n')

		// Send message to server
		_, err = fmt.Fprintf(conn, message)
		if err != nil {
			fmt.Println("Error sending message:", err.Error())
			return
		}
	}
}

func listenServer(conn net.Conn) {
	for {
		// Receive response from server
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error receiving response:", err.Error())
			return
		}
		fmt.Print(response)
	}
}
