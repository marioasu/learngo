package main

import (
	"fmt"
	// "gopkg.in/mgo.v2/bson"
)

func main() {
	// fmt.Println(bson.M{"name": "mario"})
	// fmt.Printf("%b", byte(255)) // %b 以binary格式输出 byte(n) n 超出255报溢出错误
	var a [2]int    //[0, 0]
	var b [5]string // [     ]
	var c []int     //[]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
