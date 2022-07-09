package main

import "fmt"

func main() {
	greeting("张三")
	sayHi("小李","我最好的朋友")
}
func greeting(name string) {//一个入参name，没有出参
	fmt.Println("hello" + name)
}
func sayHi(name, describe string) {//2个入参name和describe，没有出参
	fmt.Println("Hi" + name + describe)
}
