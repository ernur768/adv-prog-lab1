package structs

import (
	"net"
)

type User struct {
	name string
	conn net.Conn
}

func NewUser(conn net.Conn, name string) *User {
	return &User{
		name: name,
		conn: conn,
	}
}

func (u *User) SetName(name string) {
	name = name[0 : len(name)-2]
	u.name = name
}

func (u *User) SendMessage(message string) (int, error) {
	if len(message) > 0 {
		lastChar := message[len(message)-1]
		if lastChar != '\n' {
			message = message + "\n"
		}
	}
	return u.conn.Write([]byte(message))
}

func (u *User) CloseConnection() {
	u.conn.Close()
}
