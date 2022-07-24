package main

import "fmt"

func main() {
	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice) //[1 2 3 4 5]

	otherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, otherSlice...) //给切片追加一个切片（这里和追加元素有区别）
	fmt.Println(mySlice)                     //[1 2 3 4 5 6 7 8 9]
}
