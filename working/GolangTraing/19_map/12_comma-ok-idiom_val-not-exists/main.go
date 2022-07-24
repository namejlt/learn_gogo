package main

import "fmt"

func main() {

	myMap := map[int]string{
		0: "Golang",
		1: "Python",
		2: "Java",
		3: "PhP",
	}

	fmt.Println(myMap)

	if val, ok := myMap[7]; ok { //判断键值对是否存在
		delete(myMap, 7) //删除键为7的数据
		fmt.Println("val: ", val)
		fmt.Println("ok: ", ok)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", ok)
	}

	fmt.Println(myMap)
}
