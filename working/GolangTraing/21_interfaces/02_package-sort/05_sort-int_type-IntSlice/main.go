package main

import (
	"fmt"
	"sort"
)

func main() {
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3, 55, 84, 2}

	fmt.Println(n)
	sort.Sort(sort.IntSlice(n)) //给整形切片排序，按升序排列
	fmt.Println(n)

}
