// 从字符串中解析数字在很多程序中是一个基础常见的任务，
// 在Go 中是这样处理的。

package main

import (
	"fmt"
	"strconv" // 提供数字解析功能
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	i, _ := strconv.ParseInt("123", 0, 64) // 0表示自动推断字符串所表示的数字进制
	fmt.Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// `Atoi` 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

// [ `go run number-parsing.go` | done: 269.985135ms ]
// 	1.234
// 	123
// 	456
// 	789
// 	135
// 	strconv.Atoi: parsing "wat": invalid syntax
