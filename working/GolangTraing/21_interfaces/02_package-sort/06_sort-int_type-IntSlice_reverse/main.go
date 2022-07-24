package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n))) //sort.Reverse按照整数降序排列
	fmt.Println(n)

}
