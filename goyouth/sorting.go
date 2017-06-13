package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("String:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:      ", ints)

	// 检查一个序列是不是已经是排好序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sprted: ", s)
}

// [ `go run sorting.go` | done: 708.180433ms ]
// 	String: [a b c]
// 	Ints:       [2 4 7]
// 	Sprted:  true
