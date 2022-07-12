package main

import (
	"fmt"
)

func main() {
	var s1 []string   //声明切片但未初始化，值为nil
	var s2 [][]string //声明切片但未初始化，值为nil
	//s1[0] = "张三"
	//s1 = append(s1, "张三")
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true
}
