package main

import "fmt"

type animals struct { //type定义animals为结构体类型，包含三个字段
	dog string
	cat string
	pig string
}

func (a animals) name() string { //定义一个name方法，返回animals中的一个
	return a.pig
}

func main() {
	a1 := animals{"小狗", "小猫", "小猪"} //实例化
	a2 := animals{"大狗", "大猫", "大猪"} //实例化
	fmt.Println(a1.name())
	fmt.Println(a2.name())
}
