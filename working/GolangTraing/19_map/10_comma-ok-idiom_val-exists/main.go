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
	//delete(myMap, 2)//删除键为2的一组数据

	if val, ok := myMap[2]; ok { //判断map中某个键值对是否存在
		fmt.Println("值存在")
		fmt.Println("val", val)
		fmt.Println("ok", ok)
	} else {
		fmt.Println("不存在")
		fmt.Println("val", val)
		fmt.Println("ok", ok)
	}
	fmt.Println(myMap)
}
