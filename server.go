package main

import (
	"bufio"
	"fmt"
	"lab1/structs"
	"net"
	"strings"
)

var channels map[string]*structs.Channel

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()
	fmt.Println("Server listening on port 8080")

	channels = make(map[string]*structs.Channel)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	user := structs.NewUser(conn, conn.RemoteAddr().String())

	_, _ = user.SendMessage("commands: /join <chanel name> /create <chanel name>")

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		substrings := strings.Split(message, " ")

		if substrings[0] == "/join" {
			chanel := channels[substrings[1]]
			if chanel == nil {
				_, _ = user.SendMessage("chanel does not exists")
				continue
			}
			chanel.AddUser(user)
			return
		}
		if substrings[0] == "/create" {
			chanelName := substrings[1]
			chanel := channels[chanelName]
			if chanel != nil {
				_, _ = user.SendMessage("chanel already exists")
				continue
			}
			newChanel := structs.NewChanel(chanelName)
			channels[chanelName] = newChanel
			newChanel.AddUser(user)
			return
		}
	}
}
