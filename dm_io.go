package main

import (
	"fmt"
	"io"
	"math/rand"
)

func main() {
	fmt.Println("xx")
}

/*
type Reader interface {
    Read(p []byte) (n int, err error)
}

func ReadFull(r Reader, buf []type) (n int, err error)

crypto/rand 包里的 Reader:     var Reader io.Reader
fmt.Println(rand.Reader) // &{0 {0 0}}
Reader是一个伪随机生成器 在Linux下 Reader通过读取/dev/urandom产生的随机字符流生成随机数
*/
