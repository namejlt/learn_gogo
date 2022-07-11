package main

import "fmt"

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(5)//"/"表示取整，%表示取余
	fmt.Println(h, even)
}
