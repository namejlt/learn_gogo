package main

import "fmt"

type vehicles interface{} //定义vehicles接口

type vehicle struct { //定义vehicles结构体
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	c1 := car{} //实例化结构体car，空值
	c2 := car{}
	c3 := car{}

	p1 := plane{} //实例化结构体plane，空值
	p2 := plane{}
	p3 := plane{}

	b1 := boat{} //实例化结构体boat，空值
	b2 := boat{}
	b3 := boat{}
	rides := []vehicles{c1, c2, c3, p1, p2, p3, b1, b2, b3} //结构体car,plane,boat实现了vehicles接口

	for key, value := range rides {
		fmt.Println(key, " - ", value)
	}
}
