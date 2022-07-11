package main

import "fmt"

func main() {
	greeting := []string{"1", "2", "4", "8", "16", "32"} //定义一个字符串类型切片
	for k, v := range greeting {                         //使用for  range遍历切片
		fmt.Println(k, v)
	}
	for j := 0; j < len(greeting); j++ { //使用普通for循环遍历切片
		fmt.Println(greeting[j])
	}

}
