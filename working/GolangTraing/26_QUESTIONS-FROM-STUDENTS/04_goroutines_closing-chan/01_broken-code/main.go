package main

import "fmt"

var done chan bool

func main() {
	done = make(chan bool)
	c := fanIn(incrementor("1"), incrementor("2"))
	for n := range c {
		fmt.Println(n)
	}
}

func incrementor(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 20; i++ {
			c <- fmt.Sprintf("Process: "+s+" printing:", i)
		}
		done <- true
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	go func() {
		<-done
		<-done
		close(c)
	}()
	return c
}
