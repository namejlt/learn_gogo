package main

import "fmt"

func main() {
	switch "6" {
	case "2":
		fmt.Println("这是 2")
	case "3":
		fmt.Println("这是 3")
	case "4":
		fmt.Println("这是 4")
		fallthrough //使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	case "5":
		fmt.Println("这是 5")
		fallthrough
	case "6":
		fmt.Println("这是 6")
	case "7":
		fmt.Println("这是 7")
	}
}
