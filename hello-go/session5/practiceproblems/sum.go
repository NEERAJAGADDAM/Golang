package main

import "fmt"


func sumofdigits(n int) int {
	sum := 0
	n = abs(n) 

	for n > 0 {
		sum += n % 10 
		n /= 10       
	}
	return sum
}


func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	a := 234 
	fmt.Println("Sum of digits:", sumofdigits(a))
}
