package main

import "fmt"

var x = 0

func increment() int { //这里是个闭包。在 Golang 中，闭包是一个可以引用其作用域之外变量的函数。
	x++
	return x
}

func main() {
	//每次调用increment()，结果都不同
	fmt.Println(increment()) //1
	fmt.Println(increment()) //2
	fmt.Println(increment()) //3
	fmt.Println(increment()) //4
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
