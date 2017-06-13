// 通道方向
package main

import (
	"fmt"
)

// 定义只允许发送数据的通道ping
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 定义只允许接收数据的通道pong
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// [ `go run channel-directions.go` | done: 238.376456ms ]
// 	passed message
