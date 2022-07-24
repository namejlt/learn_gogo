package main

import (
	"encoding/json"
	"fmt"
)

type man struct { //定义一个有四个字段的结构体man
	FirstName string
	LastName  string
	Age       int
	Number    int
}

func main() {
	p1 := man{"James", "Bond", 20, 002}
	bs, _ := json.Marshal(p1) //将结构体类型转换成json字符串，Marshal返回的是字节切片
	fmt.Println(string(bs))   //将字节切片转换成字符串便于查看
}
