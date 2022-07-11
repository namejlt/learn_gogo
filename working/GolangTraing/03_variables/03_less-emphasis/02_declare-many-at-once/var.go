package main

import "fmt"

func main() {
	var message string
	var a, b, c int //同一类型的多个变量可以在一行声明，可以只写一个类型。
	a = 1

	message = "Hello World!"

	fmt.Println(message, a, b, c) //b，c声明了但没赋值，这里会打印其默认值0
}
