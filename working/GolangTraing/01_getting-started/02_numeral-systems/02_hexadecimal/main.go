package main

import "fmt"

func main() {
	fmt.Printf("%d \t %o \t %#x \n", 42, 42, 42) //42       52      0x2a
	//%d十进制，%o八进制，%x十六进制小写，%X十六进制大写，‘#’的作用是完整呈现所有数值位数

	//%#o八进制在数值前会加上数字0即052；%#x十六进制会在数值前加上0x即0x2a
	fmt.Printf("%d - %o - %x \n", 42, 42, 42)  //42 - 52 - 2a
	fmt.Printf("%d - %o - %#x \n", 42, 42, 42) //42 - 52 - 0x2a
	fmt.Printf("%d - %o - %#X \n", 42, 42, 42) //42 - 52 - 0X2A

}
