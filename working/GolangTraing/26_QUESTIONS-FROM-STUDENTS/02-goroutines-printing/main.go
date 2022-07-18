package main

import "fmt"

func main() {
	c := make(chan int) //make初始化一个通道chan
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("In goroutine!!!!!!", i)
			c <- 1
		}
		close(c)
	}()
	for n := range c {
		fmt.Println("在main中", n)
	}
}
