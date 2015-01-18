package main

import "fmt"

func main(){
    sli := make([]int, 5)
    fmt.Println(sli)
    fmt.Println(len(sli))
    fmt.Println(cap(sli))
	myMap := make(map[string] int, 20)
    fmt.Println(myMap)

    fmt.Println(len(myMap))
}

