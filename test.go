package main

import (
	"fmt"
	// "gopkg.in/mgo.v2/bson"
	"crypto/rand"
	"encoding/base64"
	"io"
	// "time"
)

func main() {
	//fmt.Println(bson.M{"name": "mario"})
	// var m = make(map[string]int)
	// fmt.Println(m)
	// m["name"] = 25
	// if a, b := m["name"]; true {
	// 	fmt.Println(a)
	// 	fmt.Println(b)
	// }
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Int())
	fmt.Println("------------")

	fmt.Println(rand.Reader) // &{0 {0 0}}
	b := make([]byte, 3)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		fmt.Println("err isnot nil")
	}

	fmt.Println(b)
	fmt.Println(base64.URLEncoding.EncodeToString(b))

	fmt.Println("------------")
}
