package main

import "fmt"

const p = "death & taxes"

func main() {
	const q = 42
	fmt.Println("p-",p)
	fmt.Println("q-",q)
	outConst(q)//调用outConst()函数
}
func outConst(q int)  {
	fmt.Println("访问q的值",q)
}
// a CONSTANT is a simple unchanging value
