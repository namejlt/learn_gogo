package main

import "fmt"

type A int //type可以对类型进行定义 A是int类型

func main() {
	var B A //B是A类型，也即是int类型
	B = 44
	fmt.Printf("%T %v\n", B, B) //main.A 44
}
