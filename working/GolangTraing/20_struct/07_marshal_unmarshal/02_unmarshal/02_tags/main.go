package main

import (
	"encoding/json"
	"fmt"
)

type man struct { //定义一个有四个字段的结构体man
	FirstName string `json:"fn"` //给结构体字段换个名称
	LastName  string `json:"ln"`
	Age       int    `json:"age"`
	Number    int    `json:"xxxx"`
}

func main() {
	var p1 man
	fmt.Println(p1.FirstName)
	fmt.Println(p1.LastName)
	fmt.Println(p1.Age)
	fmt.Println(p1.Number)

	bs := []byte(`{"fn":"James", "ln":"Bond", "age":20,"xxxx":12}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------------")
	fmt.Println(p1.FirstName)
	fmt.Println(p1.LastName)
	fmt.Println(p1.Age)
	fmt.Println(p1.Number)
	fmt.Printf("%T \n", p1)
}
