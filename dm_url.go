package main

// http.Request.Form 是一个url.Values类型 存储的是类似key=value的信息

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{} // map[string][]string结构
	v.Set("name", "s")
	v.Add("friend", "y")
	v.Add("friend", "z")
	v.Add("friend", "m")
	fmt.Println("Encode:", v.Encode()) // Encode: friend=y&friend=z&friend=m&name=s
	fmt.Println(v.Get("name"))         // s
	fmt.Println(v.Get("friend"))       // y
	fmt.Println(v.Get("friend"))       // y
	fmt.Println(v["friend"])           // [y z m]

	v2 := url.Values{"name": {"ss"}, "friends": {"tom", "jack", "Jason"}}
	fmt.Println(v2.Encode()) // friends=tom&friends=jack&friends=Jason&name=ss
	v2.Set("relatives", "y0")
	v2.Add("relatives", "yl")
	v2.Add("relatives", "yl")
	fmt.Println(v2.Get("relatives"))
	fmt.Println(v2.Encode())
	v2.Set("relatives", "just1")
	fmt.Println(v2)
}
