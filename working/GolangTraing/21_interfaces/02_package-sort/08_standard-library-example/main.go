package main

import (
	"fmt"
	"sort"
)

type person struct { //person结构体包含两个字段Name,Age
	Name string
	Age  int
}

func (p person) String() string { //String方法,并且person结构体实现了这个方法
	return fmt.Sprintf("个人信息 %s: %d", p.Name, p.Age)
}

type ByAge []person //定义ByAge为[]person类型

func (a ByAge) Len() int { //ByAge实现了 Len()方法
	return len(a)
}
func (a ByAge) Swap(i, j int) { //ByAge实现了 Swap(i, j int)方法
	a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool { //ByAge实现了 Less(i, j int)方法
	return a[i].Age < a[j].Age
}

func main() {
	p := []person{ //实例化 []person
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(p[0])   //输出 []person的第一个字段
	fmt.Println(p)      //输出 []person所有字段
	sort.Sort(ByAge(p)) //将字段按照Age升序输出信息
	fmt.Println(p)      //输出升序排列的字段

}
