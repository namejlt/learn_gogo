package main

import "fmt"

func main() {

	b := true

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	//fmt.Println(food)//无法访问局部变量food

}
