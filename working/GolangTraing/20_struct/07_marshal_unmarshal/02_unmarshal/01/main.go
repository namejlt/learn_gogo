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
	var p1 man
	fmt.Println(p1.FirstName)
	fmt.Println(p1.LastName)
	fmt.Println(p1.Age)
	fmt.Println(p1.Number)

	bs := []byte(`{"FirstName":"James", "LastName":"Bond", "Age":20, "Number":500}`)
	json.Unmarshal(bs, &p1) //因为json.UnMarshal() 函数接收的参数是字节切片,上面需要将bs转换成字节切片

	fmt.Println("--------------")
	fmt.Println(p1.FirstName)
	fmt.Println(p1.LastName)
	fmt.Println(p1.Age)
	fmt.Println(p1.Number)
	fmt.Printf("%T \n", p1)
}
