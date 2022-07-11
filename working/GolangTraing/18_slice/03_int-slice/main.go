package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3) //使用make为切片分配内存，长度为0，容量为3

	fmt.Println(mySlice)      //未对切片初始化，默认为空切片[]
	fmt.Println(len(mySlice)) //打印切片的长度0
	fmt.Println(cap(mySlice)) //打印切片的容量3

}
