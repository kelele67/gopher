// Go 提供内置的[正则表达式](http://zh.wikipedia.org/wiki/%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F)。
// 这里是 Go 中基本的正则相关功能的例子。

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们是直接使用字符串，但是对于一些其他的正则任务，
	// 你需要 `Compile` 一个优化的 `Regexp` 结构体。

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	// 这个也是查找第一次匹配的字符串的，但是返回的匹配开
	// 始和结束位置索引，而不是匹配的内容。
	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	// `Submatch` 返回完全匹配和局部匹配的字符串。例如，
	// 这里会返回 `p([a-z]+)ch` 和 `([a-z]+) 的信息
	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 带 `All` 的这个函数返回所有的匹配项，而不仅仅是首次匹配项。
	// 例如查找匹配表达式的所有项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// 这个函数提供一个正整数来限制匹配次数。

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 上面的例子中，我们使用了字符串作为参数，并使用了
	// 如 `MatchString` 这样的方法。
	// 我们也可以提供 `[]byte`参数并将 `String` 从函数名中去掉。

	fmt.Println(r.Match([]byte("peach")))

	// 创建正则表示式常量时，可以使用 `Compile` 的变体
	// `MustCompile` 。因为 `Compile` 返回两个值，不能用于常量。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// `regexp` 包也可以用来替换部分字符串为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// `Func` 变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}

// [ `go run regular-expressions.go` | done: 638.708779ms ]
// 	true
// 	true
// 	peach
// 	[0 5]
// 	[peach ea]
// 	[0 5 1 3]
// 	[peach punch pinch]
// 	[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
// 	[peach punch]
// 	true
// 	p([a-z]+)ch
// 	a <fruit>
// 	a PEACH