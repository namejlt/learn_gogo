package main

import "fmt"

func main() {

	myMap := map[string]string{
		"A": "a",
		"B": "b",
	}

	myMap["C"] = "c" //向map中添加一个k-v   C:c
	fmt.Println(myMap)
	myMap["C"] = "cc" //更新k对应的v    C:cc
	fmt.Println(myMap)
}
