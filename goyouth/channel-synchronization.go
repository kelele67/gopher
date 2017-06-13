// 通道同步
package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值通知我们完工
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
}
