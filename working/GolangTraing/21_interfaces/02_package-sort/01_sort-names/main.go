package main

import (
	"fmt"
	"sort"
)

type name []string

//声明三个方法Len() Swap(i, j int) Less(i, j int)

func (n name) Len() int {
	return len(n)
}
func (n name) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n name) Less(i, j int) bool {
	return n[i] < n[j]
}
func main() {
	studyGroup := name{"b", "t", "g", "a"}

	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

}
