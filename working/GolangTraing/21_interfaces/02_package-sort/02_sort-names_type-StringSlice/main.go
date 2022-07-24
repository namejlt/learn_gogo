package main

import (
	"fmt"
	"sort"
)

func main() {

	slice := []string{"alex", "mdddm", "dwedef", "bfgh"} //定义一个切片
	fmt.Println(slice)

	sort.StringSlice(slice).Sort()
	//sort.Sort(sort.StringSlice(slice))//对切片中的元素进行排序，按从a到z顺序
	fmt.Println(slice)
}
