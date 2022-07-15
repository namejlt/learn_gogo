package main

import "fmt"

type man struct { //定义结构体
	name string
	age  int
}

func main() {
	p1 := &man{"alex", 20} //初始化结构体指针

	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
