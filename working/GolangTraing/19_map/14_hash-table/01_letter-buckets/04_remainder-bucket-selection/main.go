package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ { //string(i)将ASCII值转换成对应的字符
		fmt.Println(i, " - ", string(i), " - ", i%12) //字符串是一个定长的字节数组
	}
}
