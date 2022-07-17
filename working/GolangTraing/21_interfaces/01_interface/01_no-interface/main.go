package main

import "fmt"

type s struct { //结构体包含一个字段
	side float64
}

func (z s) area() float64 {
	return z.side * z.side //area()返回边长的平方
}
func main() {
	ss := s{10} //实例化结构体
	fmt.Println("Area", ss.area())
}
