package main

import "fmt"

type vehicle struct { //将vehicle定义为结构体类型，三个字段
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct { //定义car结构体，三个字段，嵌套了vehicle结构体
	vehicle
	Wheels int
	Doors  int
}

type plane struct { //plane，两个字段，嵌套了vehicle结构体
	vehicle
	Jet bool
}

type boat struct { //定义car结构体，两个字段，嵌套了vehicle结构体
	vehicle
	Length int
}

func main() {
	c1 := car{} //实例化结构体car，空值
	c2 := car{}
	c3 := car{}
	cars := []car{c1, c2, c3}

	p1 := plane{} //实例化结构体plane，空值
	p2 := plane{}
	p3 := plane{}
	planes := []plane{p1, p2, p3}

	b1 := boat{} //实例化结构体boat，空值
	b2 := boat{}
	b3 := boat{}
	boats := []boat{b1, b2, b3}

	for key, value := range cars { //遍历car
		fmt.Println(key, " - ", value)
	}

	for key, value := range planes { //遍历planes
		fmt.Println(key, " - ", value)
	}

	for key, value := range boats { //遍历boats
		fmt.Println(key, " - ", value)
	}
}
