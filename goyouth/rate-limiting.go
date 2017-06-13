// 速率限制
// [_速率限制(英)_](http://en.wikipedia.org/wiki/Rate_limiting) 是
// 一个重要的控制服务资源利用和质量的途径。Go 通过 Go 协程、通
// 道和[打点器](../tickers/)优美的支持了速率限制。

package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)

	// 通过在每次请求前阻塞 `limiter` 通道的一个接收，我们限制
	// 自己每 200ms 执行一次请求
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 有时候我们想临时进行速率限制，并且不影响整体的速率控制
	// 我们可以通过[通道缓冲](channel-buffering.html)来实现。
	// 这个 `burstyLimiter` 通道用来进行 3 次临时的脉冲型速率限制。
	burstyLimiter := make(chan time.Time, 3)

	// 先将通道填充需要临时改变次的值，做好准备。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 ms 我们将添加一个新的值到 `burstyLimiter`中，
	// 直到达到 3 个的限制。
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟超过 5 个的接入请求。它们中刚开始的 3 个将
	// 由于受 `burstyLimiter` 的“脉冲”影响。
	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request_bursty", req, time.Now())
	}
}

// [ `go run rate-limiting.go` | done: 1.699068013s ]
// 	request 1 2017-06-13 23:58:46.991786311 +0800 CST
// 	request 2 2017-06-13 23:58:47.189072924 +0800 CST
// 	request 3 2017-06-13 23:58:47.388870878 +0800 CST
// 	request 4 2017-06-13 23:58:47.587049367 +0800 CST
// 	request 5 2017-06-13 23:58:47.78641758 +0800 CST
// 	request_bursty 1 2017-06-13 23:58:47.786470135 +0800 CST
// 	request_bursty 2 2017-06-13 23:58:47.78647972 +0800 CST
// 	request_bursty 3 2017-06-13 23:58:47.786484515 +0800 CST
// 	request_bursty 4 2017-06-13 23:58:47.986601407 +0800 CST
// 	request_bursty 5 2017-06-13 23:58:48.187297274 +0800 CST
