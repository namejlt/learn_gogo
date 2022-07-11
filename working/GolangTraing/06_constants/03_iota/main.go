package main

import "fmt"

const (//const专门用来定义常量
	a1 = iota//单个const声明块中从0开始取值
	a2 = iota
	a3 = iota//单个const声明块中，每增加一行声明,iota 的取值增1
)

const (
	b1 = iota//单个const声明块中，每增加一行声明,iota 的取值增1
	b2//即便声明中没有使用iota也是如此
	b3
)

func main() {
	fmt.Println(a1)//0
	fmt.Println(a2)//1
	fmt.Println(a3)//2

	fmt.Println(b1)//0
	fmt.Println(b2)//1
	fmt.Println(b3)//2
}
