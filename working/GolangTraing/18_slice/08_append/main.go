package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good"
	greeting[1] = "Bad"
	greeting[2] = "Better"

	greeting = append(greeting, "NotBetter") //append()内建函数，往切片里面添加元素并返回切片

	fmt.Println(greeting[3]) //打印添加的元素
	fmt.Println(greeting)    //打印添加元素后的切片，此时切片变为[Good Bad Better NotBetter]
}
