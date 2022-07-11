package main

import "fmt"

const (
	pi       = 3.14
	language = "Golang"
)
//以上声明相当于
//const pi      =3.14
//const language="Golang"
func main() {
	fmt.Println(pi)
	fmt.Println(language)
}
