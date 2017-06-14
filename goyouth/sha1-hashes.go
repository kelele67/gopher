// [_SHA1 散列_](http://en.wikipedia.org/wiki/SHA-1)经常用
// 生成二进制文件或者文本块的短标识。例如，[git 版本控制系统](http://git-scm.com/)
// 大量的使用 SHA1 来标识受版本控制的文件和目录。
// 这里是 Go 中如何进行 SHA1 散列计算的例子。

package main

import (
	"crypto/sha1" // Go 在多个 `crypto/*` 包中实现了一系列散列函数
	"fmt"
)

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	// SHA1 值经常以 16 进制输出，例如在 git commit 中。
	// 使用`%x` 来将散列结果格式化为 16 进制字符串。
	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}

// [ `go run sha1-hashes.go` | done: 270.279414ms ]
// 	sha1 this string
// 	cf23df2207d99a74fbe169e3eba035e633b65d94

// 你可以使用和上面相似的方式来计算其他形式的散列值。例
// 如，计算 MD5 散列，引入 `crypto/md5` 并使用 `md5.New()`
// 方法。

// 注意，如果你需要密码学上的安全散列，你需要小心的研究一下
// [哈希强度](http://en.wikipedia.org/wiki/Cryptographic_hash_function)。
