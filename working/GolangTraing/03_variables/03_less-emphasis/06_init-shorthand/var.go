package main

import "fmt"

func main() {

	// you can only do this inside a func
	message := "Hello World!"
	a, b, c := 1, false, 3
	d := 4
	e := true
	//以上变量由于首字母小写了，所以是局部变量，只能在本函数内使用，其他函数无法调用这些变量。
	fmt.Println(message, a, b, c, d, e)
}
