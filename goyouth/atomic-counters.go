// Go 中最主要的状态管理方式是通过通道间的沟通来完成的，我们
// 在[工作池](../worker-pools/)的例子中碰到过，但是还是有一
// 些其他的方法来管理状态的。这里我们将看看如何使用 `sync/atomic`
// 包在多个 Go 协程中进行 _原子计数_

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	// 我们将使用一个无符号整形数来表示（永远是正整数）这个计数器。
	var ops uint64 = 0

	// 为了模拟并发更新，我们启动 50 个 Go 协程，对计数
	// 器每隔 1ms （应为非准确时间）进行一次加一操作。
	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)

	// 为了在计数器还在被其它 Go 协程更新时，安全的使用它，
	// 我们通过 `LoadUint64` 将当前值得拷贝提取到 `opsFinal`
	// 中。和上面一样，我们需要给这个函数所取值的内存地址 `&ops`
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}

// [ `go run atomic-counters.go` | done: 1.334071682s ]
// 	ops: 5726166
