package main

import "fmt"

func main() {
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %o \t %x \n", i, i, i, i) //对1000000到1000099之间的每个数分别进行四种格式化打印：十进制。二进制，八进制，十六进制
	}
}
