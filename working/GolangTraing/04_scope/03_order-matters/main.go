package main

import "fmt"

func main() {
	//fmt.Println(x) //变量得先声明再使用
	fmt.Println(y) //全局变量y在何处都能被访问到
	//x := 42
}

var y = 42
