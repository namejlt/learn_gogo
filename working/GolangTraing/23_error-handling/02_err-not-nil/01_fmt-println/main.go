package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("1.txt") //Open打开指定文件名字的文件并读出来，否则返回err
	if err != nil {
		fmt.Println("有错误", err)
		//fmt.Println(f)
		log.Println("有错误", err)
		//log.Fatalln(err)
		//panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.
