package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	Println2(111, 1, 2, 3, 4, 5)
	fmt.Println("aaa", "cvcc", "xxx", "xxx", 1, 2, 32, 344, 3, true, false, []int{1, 2, 3})
}

func Println2(a ...int) (n int, err error) {
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	return fmt.Fprintln(os.Stdout, a)
}
