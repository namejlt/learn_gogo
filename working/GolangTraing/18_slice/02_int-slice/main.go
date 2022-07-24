package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 9, 11} //定义一个len()=cap()=6的切片
	for k, v := range xs {         //for循环，k表示切片中元素的下标，v表示下标对应的值
		fmt.Println(k, "-", v) //打印切片的key和对应的value
	}
}
