// 打点器
package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器和定时器的机制有点相似：一个通道用来发送数据。
	// 这里我们在这个通道上使用内置的 `range` 来迭代值每隔
	// 500ms 发送一次的值。
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}

// [ `go run tickers.go` | done: 1.789172116s ]
// 	Tick at 2017-06-13 23:43:09.627419716 +0800 CST
// 	Tick at 2017-06-13 23:43:10.124890273 +0800 CST
// 	Tick at 2017-06-13 23:43:10.623750228 +0800 CST
// 	Ticker stopped
