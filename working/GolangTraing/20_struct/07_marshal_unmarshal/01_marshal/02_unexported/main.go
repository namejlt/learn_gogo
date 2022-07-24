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
	p1 := man{"James", "Bond", 20, 001}
	fmt.Println(p1)
	bs, _ := json.Marshal(p1) //先转成[]byte类型
	fmt.Println(string(bs))   //再转成json字符串
}
