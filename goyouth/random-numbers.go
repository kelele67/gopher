// Go 的 `math/rand` 包提供了[伪随机数生成器（英）](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)。

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `0.0 <= f <= 1.0`。
	fmt.Println(rand.Float64())

	// `5.0 <= f <= 10.0`
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 要让伪随机数生成器有确定性，可以给它一个明确的种子。
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果使用相同的种子生成的随机数生成器，将会产生相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)

	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
}

// [ `go run random-numbers.go` | done: 895.904185ms ]
// 	81,87
// 	0.6645600532184904
// 	7.1885709359349015,7.123187485356329
// 	5,87
// 	5,87
