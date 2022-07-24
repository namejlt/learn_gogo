package main

import "fmt"

func main() {

	myMap := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
	}
	fmt.Println(myMap)
	delete(myMap, 7) //删除一个不存在的键值对，但是不会报错
	fmt.Println(myMap)
}
