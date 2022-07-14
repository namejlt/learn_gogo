package main

import "fmt"

func main() {

	myMap := map[string]string{ //map声明+赋值
		"A": "a",
		"B": "b",
		"C": "c",
		"D": "d",
	}
	fmt.Println(myMap) //打印map：map[A:a B:b C:c D:d]
	delete(myMap, "B") //删除map的key，其对应的value也会被删除
	fmt.Println(myMap) //删除后的map：map[A:a C:c D:d]
}
