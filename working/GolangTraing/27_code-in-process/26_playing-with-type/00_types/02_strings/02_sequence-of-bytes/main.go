package main

import "fmt"

func main() {
	intro := "Four score and seven years ago...."
	fmt.Println(intro)
	fmt.Println([]byte(intro)) //将intro以字节的方式输出
	// [70 111 117 114 32 115 99 111 114 101 32 97 110 100 32 115 101 118
	//101 110 32 121 101 97 114 115 32 97 103 111 46 46 46 46]
}
