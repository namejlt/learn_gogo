package main

import "fmt"

func main() {
	intro := "hello world...."
	fmt.Println(intro)
	fmt.Println(&intro) //打印intro在内存中的地址
	intro = "Hahahaha!" //重新赋值
	fmt.Println(intro)
	fmt.Println(&intro) //内存地址不变，值类型变量
}
