package main

import (
	"fmt"
)

func main() {
	// read NOTE ONE below
	c1 := incrementor("1")
	c2 := incrementor("2")

	c := fanIn(c1, c2)

	for n := range c {
		fmt.Println(n)
	}

}

func incrementor(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 20; i++ {

			c <- fmt.Sprint("Process: "+s+" printing:", i)
		}

		close(c)
	}()
	return c
}

// FAN IN
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	done := make(chan bool)

	go func(x <-chan string) {
		for n := range x {
			c <- n
		}
		done <- true
	}(input1)

	go func(x <-chan string) {
		for n := range x {
			c <- n
		}
		done <- true
	}(input2)

	go func() {
		<-done
		<-done
		close(c)
	}()

	return c
}
