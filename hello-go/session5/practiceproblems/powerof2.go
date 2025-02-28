package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number:")
	fmt.Scanln(&n) 

	if n > 0 && (n & (n-1)) == 0 { 
		fmt.Println("The given number is a power of 2")
	} else {
		fmt.Println("The given number is not a power of 2")
	}
}
