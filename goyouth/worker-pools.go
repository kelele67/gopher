// 工作池
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 为了使用 worker 工作池并且收集他们的结果，我们需要2 个通道。
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 这里启动了 3 个 worker，初始是阻塞的，因为还没有传递任务
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 这里我们发送 9 个 `jobs`，然后 `close` 这些通道来表示这些就是所有的任务了
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		<-results
	}
}

// [ `go run worker-pools.go` | done: 3.289844986s ]
// 	worker 3 processing job 1
// 	worker 1 processing job 2
// 	worker 2 processing job 3
// 	worker 3 processing job 5
// 	worker 2 processing job 4
// 	worker 1 processing job 6
// 	worker 3 processing job 7
// 	worker 2 processing job 9
// 	worker 1 processing job 8
