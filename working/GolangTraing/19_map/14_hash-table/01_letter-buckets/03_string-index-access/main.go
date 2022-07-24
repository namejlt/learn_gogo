package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0]) //取字符串"Hello"的第一个字符的ASCII值
	fmt.Println(letter)
}
