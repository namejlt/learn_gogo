package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3, 55, 84, 2}
	sort.Ints(n) //排序整形切片按升序排列
	fmt.Println(n)
}
