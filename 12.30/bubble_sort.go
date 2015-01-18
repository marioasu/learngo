package main

import (
	"fmt"
)

func main(){
	mySlice := []int{7,9,2,3,12,4,3,6}
	fmt.Println(mySlice)
	bubble_sort(mySlice)
	fmt.Println(mySlice)
}

// slice 是指针传递
func bubble_sort(slice []int){
	slice_len := len(slice)
	for i:=0; i < slice_len-1; i++ {
		for j:=0; j < slice_len-1-i; j++ {
			if slice[j] > slice[j+1] {
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}
	}
}
