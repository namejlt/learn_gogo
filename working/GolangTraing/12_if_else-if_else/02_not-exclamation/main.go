package main

import "fmt"

func main() {

	if !true { //!true==false
		fmt.Println("This did not run")
	}

	if !false { //!false==true
		fmt.Println("This ran")
	}

}
