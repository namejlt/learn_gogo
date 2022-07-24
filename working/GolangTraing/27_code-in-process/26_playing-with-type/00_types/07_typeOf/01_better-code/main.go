package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "this is stored in the variable a"
	b := 42
	c, d, e, f := 44.7, true, false, 'm'
	g := "g"
	h := `h`

	fmt.Println("a - ", reflect.TypeOf(a), " - ", a)
	fmt.Println("b - ", reflect.TypeOf(b), " - ", b)
	fmt.Println("c - ", reflect.TypeOf(c), " - ", c)
	fmt.Println("d - ", reflect.TypeOf(d), " - ", d)
	fmt.Println("e - ", reflect.TypeOf(e), " - ", e)
	fmt.Println("f - ", reflect.TypeOf(f), " - ", f)
	fmt.Println("g - ", reflect.TypeOf(g), " - ", g)
	fmt.Println("h - ", reflect.TypeOf(h), " - ", h)
	fmt.Printf("h -  %T\n", h)
}
