// 自定义排序

package main

import (
	"fmt"
	"sort"
)

type ByLength []string

// 我们在类型中实现了 `sort.Interface` 的 `Len`，`Less`
// 和 `Swap` 方法，这样我们就可以使用 `sort` 包的通用
// `Sort` 方法了，`Len` 和 `Swap` 通常在各个类型中都差
// 不多，`Less` 将控制实际的自定义排序逻辑。在我们的例
// 子中，我们想按字符串长度增加的顺序来排序，所以这里
// 使用了 `len(s[i])` 和 `len(s[j])`
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}

// [ `go run sorting-by-functions.go` | done: 728.9006ms ]
// 	[kiwi peach banana]
