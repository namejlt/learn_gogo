package main

import "fmt"

func main() {

	myMap := map[string]string{ //声明一个map，含有两个键值对
		"A": "a",
		"B": "b",
	}

	myMap["C"] = "c" //向map中增加一个键值对

	fmt.Println(len(myMap)) //打印map的长度
	fmt.Println(myMap)      //打印map的值
}
