package main

import "fmt"

func main() {

	var message = "Hello World!"
	var a, b, c = 1, 2, 3 //这里类型自动推导为int

	fmt.Println(message, a, b, c)
}
