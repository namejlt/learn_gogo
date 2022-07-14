package main

import "fmt"

func main() {
	m := hashBucket("go", 22)
	n := hashBucket("python", 12)
	fmt.Println(m)
	fmt.Println(n)
}
func hashBucket(word string, buckets int) int { //定义hashBucket函数，两个入参，一个出参
	letter := int(word[0])     //将第一个参数的第一个字符(ASCII值)转换成int，然后赋值给letter
	bucket := letter % buckets //取余
	return bucket
}
