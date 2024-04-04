package structs

import (
	"bufio"
	"fmt"
	"sync"
)

type Channel struct {
	name  string
	users []*User
	mu    sync.Mutex
}

func NewChanel(name string) *Channel {
	fmt.Println("Chanel crated: ", name)
	return &Channel{
		name:  name,
		users: make([]*User, 0, 10),
	}
}

func (c *Channel) AddUser(user *User) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.users = append(c.users, user)
	go c.HandleUser(user)
}

func (c *Channel) HandleUser(user *User) {
	defer user.CloseConnection()
	fmt.Println("New client connected:", user)

	c.SendMessageToUsers("new user connected" + user.name)

	// Read message from client
	for {
		message, err := bufio.NewReader(user.conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		c.SendMessageFromUser(user, message)
	}
}

func (c *Channel) SendMessageFromUser(user *User, message string) {
	message = user.name + ": " + message
	for _, chanelUser := range c.users {
		if chanelUser == user {
			continue
		}
		_, err := chanelUser.SendMessage(message)
		if err != nil {
			fmt.Println("Send message. Error writing:", err.Error())
			return
		}
	}
}

func (c *Channel) SendMessageToUsers(message string) {
	for _, user := range c.users {
		_, err := user.SendMessage(message)
		if err != nil {
			fmt.Println("Send message. Error writing:", err.Error())
			return
		}
	}
}
