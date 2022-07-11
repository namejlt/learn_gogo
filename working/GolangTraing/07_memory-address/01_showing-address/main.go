package main

import "fmt"

func main() {
	a := 40
	fmt.Println("a-", &a)
	fmt.Println("a's memory address-", &a)    //变量a在内存中的地址是0xc0000b6010
	fmt.Printf("%d \n", &a)                //十六进制地址以十进制格式化输出824634466320
}
