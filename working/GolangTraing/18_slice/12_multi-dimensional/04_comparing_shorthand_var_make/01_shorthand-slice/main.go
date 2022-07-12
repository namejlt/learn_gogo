package main

import (
	"fmt"
)

func main() {
	s1 := []string{}   //初始化为[]空切片
	s2 := [][]string{} //初始化为[]空切片
	//s1[0] = "张三"//空切片无法赋值
	//s1 = append(s1, "张三")//空切片可以追加元素
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(s1 == nil) //false
	fmt.Println(s2 == nil) //false
}
