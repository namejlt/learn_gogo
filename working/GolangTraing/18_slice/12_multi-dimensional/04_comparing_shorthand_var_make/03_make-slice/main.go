package main

import (
	"fmt"
)

func main() {
	s1 := make([]string, 5) //初始化并分配内存空间
	s2 := make([][]string, 5)
	fmt.Println(s1)
	fmt.Println(s2)
	s1[0] = "张三"
	s1 = append(s1, "张三")
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(s1 == nil) //true
	fmt.Println(s2 == nil) //true
}
