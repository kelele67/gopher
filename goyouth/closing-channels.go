package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")

				done <- true
				return
			}
		}
	}()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

// [ `go run closing-channels.go` | done: 266.719026ms ]
// 	sent job 0
// 	sent job 1
// 	sent job 2
// 	received job 0
// 	sent job 3
// 	sent all jobs
// 	received job 1
// 	received job 2
// 	received job 3
// 	received all jobs
