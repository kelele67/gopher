package main

import (
	"fmt"
)

func Factorial(x int) (result int) {
	if x == 0 {
		result = 1
	} else {
		result = x * Factorial(x-1)
	}
	return
}

func main() {
	i := 15
	fmt.Printf("%d 阶乘 是 %d\n", i, Factorial(i))
}
