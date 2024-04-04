package structs

import (
	"bufio"
	"fmt"
	"sync"
)

type Channel struct {
	name    string
	users   []*User
	history []string
	mu      sync.Mutex
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
	go c.handleUser(user)
}

func (c *Channel) handleUser(user *User) {
	defer user.CloseConnection()
	c.SendMessageToUsers("new user connected: " + user.name)

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
	c.addToHistory(message)
}

func (c *Channel) SendMessageToUsers(message string) {
	for _, user := range c.users {
		_, err := user.SendMessage(message)
		if err != nil {
			fmt.Println("Send message. Error writing:", err.Error())
			return
		}
	}
	c.addToHistory(message)
}

func (c *Channel) addToHistory(message string) {
	c.history = append(c.history, message)
}

func (c *Channel) GetHistory() []string {
	return c.history
}
