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
*/
