package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println("这里可以访问到x值为：", x)
		y := "the credit belongs with the one who is in the ring"
		fmt.Println(y)
	}
	//fmt.Println(y) // 这里访问不到变量y的值
}
