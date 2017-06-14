package main

import (
	"fmt"
	"time"
)

func main() {
	// 分别使用带 `Unix` 或者 `UnixNano` 的 `time.Now`
	// 来获取从自[协调世界时](http://zh.wikipedia.org/wiki/%E5%8D%94%E8%AA%BF%E4%B8%96%E7%95%8C%E6%99%82)
	// 起到现在的秒数或者纳秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意 `UnixMillis` 是不存在的，所以要得到毫秒数的话，
	// 你要自己手动的从纳秒转化一下。
	millis := nanos / 1000000

	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 你也可以将协调世界时起的整数秒或者纳秒转化到相应的时间。
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

// [ `go run epoch.go` | done: 632.739604ms ]
// 	2017-06-14 13:03:20.124520477 +0800 CST
// 	1497416600
// 	1497416600124
// 	1497416600124520477
// 	2017-06-14 13:03:20 +0800 CST
// 	2017-06-14 13:03:20.124520477 +0800 CST
