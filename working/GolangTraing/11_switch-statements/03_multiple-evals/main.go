package main

import "fmt"

func main() {
	switch "12" {
	case "11", "12": //case后面可以添加多个匹配对象
		fmt.Println("OK1")
	case "14", "15":
		fmt.Println("OK2")
	case "17", "18":
		fmt.Println("OK3")
	}
}
