package main

import (
	"fmt"
)

type player struct {
	FirstName string
	LastName  string
	Age       int
}
type star struct {
	player
	FirstName string
	Rich      bool
}

func main() {
	p1 := star{
		player: player{
			FirstName: "三",
			LastName:  "张",
			Age:       36,
		},
		FirstName: "四",
		Rich:      false,
	}
	p2 := star{
		player: player{
			FirstName: "五",
			LastName:  "王",
			Age:       55,
		},
		FirstName: "六",
		Rich:      true,
	}
	fmt.Println(p1.FirstName, p1.player.FirstName) //注意访问成员变量的先后顺序
	fmt.Println(p2.FirstName, p2.player.FirstName)
}
