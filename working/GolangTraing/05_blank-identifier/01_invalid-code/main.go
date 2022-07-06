package main

import "fmt"

func main() {
	a:="stored in a"
	//b:="stored in b"
	fmt.Println("a - ",a)//报错：b declared but not used
	//go语言中变量声明后一定要使用
}
