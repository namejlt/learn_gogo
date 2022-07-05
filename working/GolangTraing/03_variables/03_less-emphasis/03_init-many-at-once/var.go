package main

import "fmt"

func main() {

	var message = "Hello World!"
	var a, b, c int = 1, 2, 3 //这里类型int可以省略，因为go语言带有类型推导的特点

	fmt.Println(message, a, b, c)
}
