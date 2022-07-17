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

func (c circle) area() float64 { //circle结构体也实现了area方法，则实现了shape接口
	return math.Pi * c.radius * c.radius //返回圆的面积
}
func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// a new method which takes the INTERFACE TYPE shape
func totalArea(shapes ...shape) float64 { //传入实例化后的接口
	var area float64
	for _, s := range shapes { //循环接口中的实例并取和，用area表示
		area += s.area()
	}
	return area
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
	fmt.Println("Total Area: ", totalArea(c, s)) //c,s表示接口实例化
}
