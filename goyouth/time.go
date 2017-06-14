// Go 对时间和时间段提供了大量的支持；这里是一些例子。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	// 这些方法来比较两个时间，分别测试一下是否是之前，
	// 之后或者是同一时刻，精确到秒。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// 方法 `Sub` 返回一个 `Duration` 来表示两个时间点的间隔时间。
	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 你可以使用 `Add` 将时间后移一个时间间隔，
	// 或者使用一个 `-` 来将时间前移一个时间间隔。
	p(then.Add(diff))
	p(then.Add(-diff))
}

// [ `go run time.go` | done: 669.902007ms ]
// 	2017-06-14 12:41:52.354118446 +0800 CST
// 	2009-11-17 20:34:58.651387237 +0000 UTC
// 	2009
// 	November
// 	17
// 	20
// 	34
// 	58
// 	651387237
// 	UTC
// 	Tuesday
// 	true
// 	false
// 	false
// 	66368h6m53.702731209s
// 	66368.11491742534
// 	3.9820868950455203e+06
// 	2.3892521370273122e+08
// 	238925213702731209
// 	2017-06-14 04:41:52.354118446 +0000 UTC
// 	2002-04-23 12:28:04.948656028 +0000 UTC
