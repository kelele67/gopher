// Go 支持通过基于描述模板的时间格式化和解析。

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(time.RFC3339, "2017-06-14T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))

	// 来指定给定时间/字符串的格式化/解析方式。时间一定要按照
	// 如下所示：2006为年，15 为小时，Monday 代表星期几，等等。
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 对于纯数字表示的时间，你也可以使用标准的格式化字
	// 符串来提出出时间值得组成。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// `Parse` 函数在输入的时间格式不正确是会返回一个
	// 错误。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}

// [ `go run time-formatting-parsing.go` | done: 737.795742ms ]
// 	2017-06-14T13:40:49+08:00
// 	2017-06-14 22:08:41 +0000 +0000
// 	1:40PM
// 	Wed Jun 14 13:40:49 2017
// 	2017-06-14T13:40:49.935919+08:00
// 	0000-01-01 20:41:00 +0000 UTC
// 	2017-06-14T13:40:49-00:00
// 	parsing time "8:41PM" as "Mon Jan _2 15:04:05 2006": cannot parse "8:41PM" as "Mon"
