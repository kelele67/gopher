package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var file1 = "/Users/Kelele67/Desktop/gopher/goyouth/data1"
var file2 = "/Users/Kelele67/Desktop/gopher/goyouth/data2"

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(file1, d1, 0644)
	check(err)

	// 对于更细粒度的写入，先打开一个文件。
	f, err := os.Create(file2)
	check(err)
	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 `Sync` 来将缓冲区的信息写入磁盘。
	f.Sync()

	// `bufio` 提供了和我们前面看到的带缓冲的读取器一
	// 样的带缓冲的写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用 `Flush` 来确保所有缓存的操作已写入底层写入器。
	w.Flush()
}

// [ `go run writing-files.go` | done: 310.211341ms ]
// 	wrote 5 bytes
// 	wrote 7 bytes
// 	wrote 9 bytes
