package main

import "fmt"

func main() {

	if !true { //!true==false
		fmt.Println("This did not run")
	}

	if !false { //!false==true
		fmt.Println("This ran")
	}
	if !false == true {
		fmt.Println("!false和true等价")
	}
	if !true == false {
		fmt.Println("!true==false等价")
	}

}
