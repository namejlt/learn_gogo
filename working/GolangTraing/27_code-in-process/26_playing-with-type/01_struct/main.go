package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{"James", 20} //实例化结构体person
	fmt.Println(p1)
	fmt.Println(p1.name) //通过.的方式取结构体里面的字段
	fmt.Println(p1.age)
	fmt.Printf("%T\n", p1)
}
