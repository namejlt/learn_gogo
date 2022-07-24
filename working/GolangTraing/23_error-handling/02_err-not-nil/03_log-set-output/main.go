package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	f, err := os.Create("log.txt") //
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(f)
}

func main() {
	_, err := os.Open("log.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		log.Println("发生错误", err)
		//		log.Fatalln(err)
		//		panic(err)
	}
}

/*
Package log implements a simple logging package ... writes to standard error and prints the date and time of each logged message ... the Fatal functions call os.Exit(1) after writing the log message ... the Panic functions call panic after writing the log message.
*/

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
