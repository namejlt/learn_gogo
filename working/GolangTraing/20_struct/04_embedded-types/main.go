package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person //结构体嵌套,这里把person结构体当做doubleZero结构体的成员
	man    bool
}

func main() {
	p1 := doubleZero{ //结构体实例化为p1
		person: person{ //成员1
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		man: true, //成员2
	}

	p2 := doubleZero{ //结构体实例化为p2
		person: person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		man: false,
	}

	fmt.Println(p1.person.First, p1.person.Last, p1.person.Age, p1.man)
	fmt.Println(p2.First, p2.Last, p2.Age, p2.man)
}
