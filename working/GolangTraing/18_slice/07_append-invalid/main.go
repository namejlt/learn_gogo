package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good"
	greeting[1] = "Bad"
	greeting[2] = "Better"
	greeting[3] = "NotBetter"

	fmt.Println(greeting[2]) //由于切片容量为3，而索引有四个，因此将返回索引超出范围并报错
}
