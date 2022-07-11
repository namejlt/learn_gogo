package main

import "fmt"

func main() {
	greeting("张三")
	sayHi("小李", "我最好的朋友")
}
func greeting(name string) string { //一个入参name，return一个string类型
	return fmt.Sprint("hello" + name)
}
func sayHi(name, describe string) (s string) { //2个入参name和describe，return一个string类型的s
	s = fmt.Sprint("Hi" + name + describe)
	return
}

// greet is declared with a parameter
// when calling greet, pass in an argument
