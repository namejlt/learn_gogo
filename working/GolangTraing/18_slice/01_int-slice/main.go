package main

import "fmt"

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11} //定义一个len()=cap()=6的切片
	fmt.Println(mySlice)
	fmt.Printf("%T", mySlice)
	fmt.Println(len(mySlice)) //打印切片的长度
	fmt.Println(cap(mySlice)) //打印切片的容量
}
