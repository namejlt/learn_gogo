package main

import "fmt"

func main() {
	var x int = 5
	str := "Hello world " + fmt.Sprint(x) //Sprint(x)函数将x类型以字符串的形式输出
	fmt.Println(str)
}
