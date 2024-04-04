package structs

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type App struct {
	channels map[string]*Channel
}

func NewApp() *App {
	return &App{
		channels: make(map[string]*Channel),
	}
}

func (a *App) Run() {
	go a.listen()
	a.HandleServerCommands()
}

func (a *App) listen() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		go a.handleConnection(conn)
	}
}

func (a *App) handleConnection(conn net.Conn) {
	user := NewUser(conn, conn.RemoteAddr().String())
	_, _ = user.SendMessage("enter your name:")
	name, err := bufio.NewReader(user.conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
		return
	}
	user.SetName(name)

	_, _ = user.SendMessage("commands: /join <chanel name> /create <chanel name>")

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		substrings := strings.Split(message, " ")
		command := substrings[0]

		if command == "/join" {
			channelName := substrings[1]
			chanel := a.channels[channelName]
			if chanel == nil {
				_, _ = user.SendMessage("chanel does not exists")
				continue
			}
			chanel.AddUser(user)
			return
		}
		if command == "/create" {
			chanelName := substrings[1]
			chanel := a.channels[chanelName]
			if chanel != nil {
				_, _ = user.SendMessage("chanel already exists")
				continue
			}
			newChanel := NewChanel(chanelName)
			a.channels[chanelName] = newChanel
			newChanel.AddUser(user)
			return
		}
	}
}

func (a *App) HandleServerCommands() {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, _ := reader.ReadString('\n')
		substrings := strings.Split(message, " ")
		command := substrings[0]

		if command == "/history" {
			channelName := substrings[1]
			a.showHistory(channelName)
		}
		if command == "/channels" {
			a.showChannels()
		}
	}
}

func (a *App) showHistory(channelName string) {
	channel := a.channels[channelName]
	if channel == nil {
		fmt.Println("channel does not exists")
		return
	}
	fmt.Printf("\nchanel %s messages: ---------\n", channelName)
	for _, message := range channel.GetHistory() {
		if message == "" || message == "\n" {
			continue
		}
		fmt.Println(message)
	}
	fmt.Println("---------")
}

func (a *App) showChannels() {
	fmt.Println("\nchannels: ---------")
	for name := range a.channels {
		fmt.Print(name)
	}
	fmt.Println("---------")
}
