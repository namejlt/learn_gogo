package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"a", "b", "c", "d"}

	fmt.Println(s)                               //[a b c d]
	sort.Sort(sort.Reverse(sort.StringSlice(s))) //sort.Reverse按照a-z相反的顺序排列
	fmt.Println(s)                               //[d c b a]
}
