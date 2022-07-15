package main

import "fmt"

type A int

func main() {
	var B A //B是A类型，也即是int类型
	B = 44
	fmt.Printf("%T %v\n", B, B) //main.A 44
	var C int
	C = 25
	fmt.Printf("%T %v\n", C, C) //int 25

}
