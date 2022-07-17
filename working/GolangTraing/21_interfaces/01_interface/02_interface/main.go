package main

import "fmt"

type square struct { //定义一个结构体square
	side float64
}

func (z square) area() float64 { //结构体square实现了shape接口的方法area（），则实现了shape接口
	return z.side * z.side
}

type shape interface { //定义一个shape接口类型，包含一个area（）方法
	area() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10} //结构体square实例化
	fmt.Printf("%T\n", s)
	info(s) //调用info函数，输出结构体square的值和area()方法的值
}
