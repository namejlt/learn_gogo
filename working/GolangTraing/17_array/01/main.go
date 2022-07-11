package main

import "fmt"

func main() {
	var x [58]int
	fmt.Println(x)//声明数组后未赋值，默认元素值为0
	fmt.Println(len(x))//内建函数len()返回数组长度
	fmt.Println(x[42])//切片中第43个数的大小
	x[42] = 777//把777赋值给数组第43个元素
	fmt.Println(x[42])
}
