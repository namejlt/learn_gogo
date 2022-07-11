package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 输出a的值：43
	fmt.Println(&a) // 输出a在内存中的地址：0x20818a220

	var b = &a      //将a的地址值赋值给b
	fmt.Println(b)  // b的值就是a的地址值：0x20818a220
	fmt.Println(*b) // 指针变量b的值为：43

	*b = 42         //指针类型是一种引用类型
	                // b says, "The value at this address, change it to 42"
	fmt.Println(a)  // 42

}
