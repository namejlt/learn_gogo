package main

import "fmt"

func main() {
	for i := 32; i <= 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i) //从32到122格式化输出每个数字对应的四种格式
	} //%d十进制表示，%b二进制表示，%x十六进制表示，%q对应ASCII 编码
}
