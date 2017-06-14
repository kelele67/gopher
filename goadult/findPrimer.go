// Go 的并发属于 CSP 并发模型的一种实现，
// CSP 并发模型的核心概念是：“不要通过共享内存来通信，而应该通过通信来共享内存”。
// 这在 Go 语言中的实现就是 Goroutine 和 Channel。
// 在1978发表的 CSP 论文中有一段使用 CSP 思路解决问题的描述。
// “Problem: To print in ascending order all primes less than 10000.
// Use an array of processes, SIEVE, in which each process inputs a prime from its predecessor and prints it.
// The process then inputs an ascending stream of numbers from its predecessor and passes them on to its successor,
// suppressing any that are multiples of the original prime.”
// 要找出10000以内所有的素数，这里使用的方法是筛法，
// 即从2开始每找到一个素数就标记所有能被该素数整除的所有数。
// 直到没有可标记的数，剩下的就都是素数。下面以找出10000以内所有素数为例，
// 借用 CSP 方式解决这个问题。

package main

import (
	"fmt"
)

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}

func main() {
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 10000; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}
