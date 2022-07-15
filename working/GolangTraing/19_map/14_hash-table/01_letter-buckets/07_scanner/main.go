package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	const input = "what day it is today?"
	const a = "what"
	scanner := bufio.NewScanner(strings.NewReader(input)) //NewReader(s string)函数的入参是一个string字符串，返回指针*Reader
	fmt.Println(strings.NewReader(input))                 //&{what day it is today? 0 -1}
	fmt.Println(scanner)                                  //&{0xc000064020 0x465960 65536 [] [] 0 0 <nil> 0 false false}
	scanner.Split(bufio.ScanWords)                        //
	fmt.Println(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
