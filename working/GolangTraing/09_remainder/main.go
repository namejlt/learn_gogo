package main

import "fmt"

func main() {
	x := 13 % 3 // "%"为取余运算符，4...1，值为1
	y := 13 / 3 // "/"为取整运算符，4...1，值为4
	fmt.Println(y)
	fmt.Println(x)
	if x == 1 {//x==1返回值为true，故执行fmt.Println("Odd")，打印值为Odd
		fmt.Println("Odd")
	} else {
		fmt.Println("Even")
	}

	for i := 1; i < 70; i++ {
		if i%2 == 1 {//i取1~69，当表达式i%2 == 1为true时，打印Odd,否则打印Even。
			fmt.Println("Odd")// i为质数
		} else {
			fmt.Println("Even")//i为偶数
		}
	}
}
