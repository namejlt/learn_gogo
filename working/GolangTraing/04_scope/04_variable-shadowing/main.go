package main

import "fmt"

func max(x int) int {
	return 42 + x
}

func main() {
	//max := max(7)
	//fmt.Println(max) // max is now the result, not the function
	//上面max既是整形变量又是函数名称，难以区分，不建议这样做
	// don't do this; bad coding practice to shadow variables

	result := max(9)
	fmt.Println(result)
}
