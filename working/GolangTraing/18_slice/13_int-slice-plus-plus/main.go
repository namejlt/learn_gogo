package main

import "fmt"

func main() {

	mySlice := make([]int, 1)
	fmt.Println(mySlice[0]) //空切片的默认元素值为0
	mySlice[0] = 7          //给切片第一个元素赋值
	fmt.Println(mySlice[0])
	mySlice[0]++ //第一个元素自加1
	fmt.Println(mySlice[0])
}
