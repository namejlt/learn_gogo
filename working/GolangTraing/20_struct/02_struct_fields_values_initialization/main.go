package main

import "fmt"

type animals struct { //type定义animals为结构体类型，包含三个字段
	dog string
	cat string
	pig string
}

func main() {
	a1 := animals{ //结构体实例化为a1
		dog: "小狗",
		cat: "小猫",
		pig: "小猪",
	}
	a2 := animals{ ////结构体实例化为a2
		dog: "大狗",
		cat: "大猫",
		pig: "大猪",
	}
	fmt.Println(a1, a2)
}
