package main

import "fmt"

var x = 42 //全局标量，包内可用

func main() {
	fmt.Println(x) //调用变量x并打印出x的值
	foo()          //调用foo函数
}

func foo() { //foo函数的作用输出x的值
	fmt.Println(x)
}
