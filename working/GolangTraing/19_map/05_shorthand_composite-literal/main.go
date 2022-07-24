package main

import "fmt"

func main() {

	myGlang := map[string]string{ //声明一个map，含有3个键值对
		"golang": "1001",
		"python": "1002",
		"php":    "1002",
	}
	fmt.Println(myGlang)         //打印map的值：map[golang:1001 php:1002 python:1002]
	for s, s2 := range myGlang { //遍历map并以k:v的形式输出
		fmt.Println(s + ":" + s2)
	}

}
