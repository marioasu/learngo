package main

import (
	"fmt"
	"time"
)

func main() {
	/* 类型
	   time.Now() --> Time
	   time.Unix(sec int64, nsec int64) --> Time
	   Time.Unix() --> int64
	   Time.UnixNano() --> int64
	   Time.Weekday() --> Weekday
	*/

	// 获得时间戳
	t := time.Now().Unix()
	fmt.Println(t)

	// 时间戳转字符串
	fmt.Println(time.Unix(t, 0).String())

	//格式化
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Weekday().String())
}
