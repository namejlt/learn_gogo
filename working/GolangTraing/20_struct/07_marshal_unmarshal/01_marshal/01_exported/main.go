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
	p1 := man{"alex", "shao", 20, 007} //结构体实例化
	p2, _ := json.Marshal(p1)          //p2是[]byte类型
	fmt.Println("p2", p2)
	fmt.Printf("%T \n", p2)
	fmt.Println(string(p2)) //将[]byte类型转换成string类型
}
