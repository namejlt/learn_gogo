package main

import (
	"fmt"
	"math"
)

type square struct { //结构体sqaure
	side float64
}

// another shape
type circle struct { //结构体circle
	radius float64
}

type shape interface { //shape接口包含一个area方法
	area() float64
}

func (s square) area() float64 { //square结构体实现了area方法，则实现了shape接口
	return s.side * s.side //返回边长的平方
}

// which implements the shape interface
func (c circle) area() float64 { //circle结构体也实现了area方法，则实现了shape接口
	return math.Pi * c.radius * c.radius //返回圆的面积
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10} //实例化
	c := circle{5}  //实例化
	info(s)
	info(c)
}
