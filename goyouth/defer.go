// _Defer_ 被用来确保一个函数调用在程序执行结束前执行。同
// 样用来执行一些清理工作。 `defer` 用在像其他语言中的
// `ensure` 和 `finally`用到的地方

package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/Users/Kelele67/Desktop/gopher/goyouth/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

// 执行程序，确认这个文件在写入后是已关闭的。

// [ `go run defer.go` | done: 388.292458ms ]
// 	creating
// 	Writing
// 	closing
