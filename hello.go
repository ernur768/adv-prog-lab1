package main

import "fmt"

func main() {
	message := "hello\n"

	if len(message) > 0 {
		lastChar := message[len(message)-1]
		if lastChar != '\n' {
			fmt.Println("does not end with new line")
		}
	}
}
