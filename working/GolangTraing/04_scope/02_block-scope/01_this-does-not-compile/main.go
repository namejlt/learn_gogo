package main

import "fmt"

func main() {
	x := 42 //局部变量
	fmt.Println(x)
	foo()
}

func foo() {
	// no access to x
	// this does not compile
	//fmt.Println(x)
	//由于x是局部变量，在main（）函数外无法访问。
}
