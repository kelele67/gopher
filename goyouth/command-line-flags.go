// [_命令行标志_](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// 是命令行程序指定选项的常用方式。例如，在 `wc -l` 中，
// 这个 `-l` 就是一个命令行标志。

package main

import (
	"flag" // `flag` 包，支持基本的命令行标志解析。
	"fmt"
)

func main() {
	// 基本的标记声明仅支持字符串、整数和布尔值选项。
	// 这里我们声明一个默认值为 `"foo"` 的字符串标志 `word`
	// 并带有一个简短的描述。这里的 `flag.String` 函数返回一个字
	// 符串指针（不是一个字符串值），在下面我们会看到是如何
	// 使用这个指针的。
	wordPtr := flag.String("word", "foo", "astring")

	// 使用和声明 `word` 标志相同的方法来声明 `numb` 和 `fork` 标志。
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。注
	// 意在标志声明函数中需要使用该参数的指针
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有标志都声明完成以后，调用 `flag.Parse()` 来执行命令行解析。
	flag.Parse()

	// 这里我们将仅输出解析的选项以及后面的位置参数。
	// 注意，我们需要使用类似 `*wordPtr` 这样的语法来对指针解引用，
	// 从而得到选项的实际值。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("bool:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// ┌─(~/Desktop/gopher/goyouth)───────────────(kelele67@liuqideMacBook-Pro:s001)─┐
// └─(15:29:10 on master ✹ ✭)──> ./command-line-flags                ──(三, 614)─┘
// word: foo
// numb: 42
// bool: false
// svar: bar
// tail: []
// ┌─(~/Desktop/gopher/goyouth)───────────────(kelele67@liuqideMacBook-Pro:s001)─┐
// └─(15:39:13 on master ✹ ✭)──> ./command-line-flags -word=opt      ──(三, 614)─┘
// word: opt
// numb: 42
// bool: false
// svar: bar
// tail: []
// ┌─(~/Desktop/gopher/goyouth)───────────────(kelele67@liuqideMacBook-Pro:s001)─┐
// └─(15:39:40 on master ✹ ✭)──> ./command-line-flags -word=opt a1 a2 a3
// word: opt
// numb: 42
// bool: false
// svar: bar
// tail: [a1 a2 a3]
// ┌─(~/Desktop/gopher/goyouth)───────────────(kelele67@liuqideMacBook-Pro:s001)─┐
// └─(15:39:53 on master ✹ ✭)──> ./command-line-flags -word=opt a1 a2 a3 -numb=7
// word: opt
// numb: 42
// bool: false
// svar: bar
// tail: [a1 a2 a3 -numb=7]
// ┌─(~/Desktop/gopher/goyouth)───────────────(kelele67@liuqideMacBook-Pro:s001)─┐
// └─(15:40:12 on master ✹ ✭)──> ./command-line-flags -h             ──(三, 614)─┘
// Usage of ./command-line-flags:
//   -fork
//     	a bool
//   -numb int
//     	an int (default 42)
//   -svar string
//     	a string var (default "bar")
//   -word string
//     	astring (default "foo")
// ┌─(~/Desktop/gopher/goyouth)───────────────(kelele67@liuqideMacBook-Pro:s001)─┐
// └─(15:40:24 on master ✹ ✭)──> ./command-line-flags -wat       2 ↵ ──(三, 614)─┘
// flag provided but not defined: -wat
// Usage of ./command-line-flags:
//   -fork
//     	a bool
//   -numb int
//     	an int (default 42)
//   -svar string
//     	a string var (default "bar")
//   -word string
//     	astring (default "foo")
