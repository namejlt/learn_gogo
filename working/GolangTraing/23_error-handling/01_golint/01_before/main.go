package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)//调用函数evalInt
	fmt.Println(str)
}

func evalInt(n int) string {//函数evalInt比较x和10的大小
	if n > 10 {
		return fmt.Sprintf("x比10大")
	} else {
		return fmt.Sprintf("x比10小")
	}
}
