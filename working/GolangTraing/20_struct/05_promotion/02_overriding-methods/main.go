package main

import (
	"fmt"
)

type player struct {
	Name string
	Age  int
}

type people struct {
	player
	man bool
}

func (p player) Greeting() {
	fmt.Println("a busketball  player")
}

func (p2 people) Greeting() {
	fmt.Println("hello")
}

func main() {
	p1 := player{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := people{
		player: player{
			Name: "alex",
			Age:  23,
		},
		man: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.player.Greeting()
}
