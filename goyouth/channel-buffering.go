// 通道缓冲
package main

import (
	"fmt"
)

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	// 有缓冲区，所以没有对应的并发接收方 也能发送这些值
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
