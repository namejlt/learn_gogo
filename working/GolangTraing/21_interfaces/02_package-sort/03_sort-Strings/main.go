package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	s1 := []string{"z", "j", "a", "g"}
	sort.Strings(s) //s切片中的元素在内存中的位置已经发生变化，按顺序排列
	sort.Strings(s1)
	fmt.Println(s)
	fmt.Println(s1)
}
