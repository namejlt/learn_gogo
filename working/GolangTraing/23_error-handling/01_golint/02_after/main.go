package main

import "fmt"

func main() {
	x := 1
	str := evalInt(x)
	fmt.Println(str)
}

func evalInt(n int) string {
	if n > 10 {
		return fmt.Sprint("x比10大")
	}//这里去掉了else,结果来说一样

	return fmt.Sprint("x比10小")
}
