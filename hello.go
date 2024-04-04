package main

import "fmt"

func main() {
	// Byte 13 is carriage return character
	byteValue := byte(13)

	// Convert byte to string
	stringValue := string(byteValue)

	// Print the result
	fmt.Printf("Byte 13 as string:%s0", stringValue)
}
