package main

import "fmt"

func main() {
	a := 10
	b := "golang"
	c := 4.17
	d := true
	e := "Hello"
	f := `Do you like me hat?`
	g := 'M'
	//%v:输出值的默认格式
	//%T：输出值相对应的类型
	fmt.Printf("%v %T \n", a, a) //10 int
	fmt.Printf("%v %T \n", b, b) //golang string
	fmt.Printf("%v %T \n", c, c) //4.17 float64
	fmt.Printf("%v %T \n", d, d) //true bool
	fmt.Printf("%v %T \n", e, e) //Hello string
	fmt.Printf("%v %T \n", f, f) //Do you like me hat? string
	fmt.Printf("%v %T \n", g, g) //77 int32（rune）
}
