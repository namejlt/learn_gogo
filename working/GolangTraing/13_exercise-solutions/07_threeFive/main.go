package main

import "fmt"

func main() {
	c := 0
	for i := 0; i < 15; i++ {
		if i%3 == 0 { //3,6,9,12
			c += i
		} else if i%5 == 0 { //5,10
			c += i
		}
	}
	fmt.Println(c) //3+6+9+12+5+10=45
}

/*

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/
