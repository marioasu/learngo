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

/*
source code
type Time struct {
	// sec gives the number of seconds elapsed since
	// January 1, year 1 00:00:00 UTC.
	sec int64

	// nsec specifies a non-negative nanosecond
	// offset within the second named by Seconds.
	// It must be in the range [0, 999999999].
	//
	// It is declared as uintptr instead of int32 or uint32
	// to avoid garbage collector aliasing in the case where
	// on a 64-bit system the int32 or uint32 field is written
	// over the low half of a pointer, creating another pointer.
	// TODO(rsc): When the garbage collector is completely
	// precise, change back to int32.
	nsec uintptr

	// loc specifies the Location that should be used to
	// determine the minute, hour, month, day, and year
	// that correspond to this Time.
	// Only the zero Time has a nil Location.
	// In that case it is interpreted to mean UTC.
	loc *Location
}

*/
