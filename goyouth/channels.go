// 通道
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

// [ `go run channels.go` | done: 568.211781ms ]
// 	ping
