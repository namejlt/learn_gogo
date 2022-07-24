package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	a1 := dog{animal{"woof"}, true}
	a2 := cat{animal{"meow"}, true}
	specs(a1)
	specs(a2)
}
