// 在前面的例子中，我们看到了如何使用原子操作来管理简单的计数器。
// 对于更加复杂的情况，我们可以使用一个_[互斥锁](http://zh.wikipedia.org/wiki/%E4%BA%92%E6%96%A5%E9%94%81)_
// 来在 Go 协程间安全的访问数据。

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0

	// 这里我们运行 100 个 Go 协程来重复读取 state
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 为了确保这个 Go 协程不会再调度中饿死，我们
				// 在每次操作后明确的使用 `runtime.Gosched()`
				// 进行释放。这个释放一般是自动处理的，像例如
				// 每个通道操作后或者 `time.Sleep` 的阻塞调用后
				// 相似，但是在这个例子中我们需要手动的处理。
				runtime.Gosched()
			}
		}()
	}

	// 同样的，我们运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	// 让这 10 个 Go 协程对 `state` 和 `mutex` 的操作运行 1 s
	time.Sleep(time.Second)

	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 对 `state` 使用一个最终的锁，显示它是如何结束的。
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

// [ `go run mutexes.go` | done: 1.59620452s ]
// 	ops: 3971161
// 	state: map[4:97 3:19 0:50 2:91 1:0]
