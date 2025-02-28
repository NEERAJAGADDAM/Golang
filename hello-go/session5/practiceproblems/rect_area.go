package main

import "fmt"

type RectArea struct {
	
	Length float64
	Width float64
	RectangularArea float64
}


func main() {
	
	var Length float64 
	fmt.Println("Length:")
	fmt.Scanln(&Length)

	var Width float64 
	fmt.Println("Width:")
	fmt.Scanln(&Width)
 
	var RectangularArea float64

	RectangularArea = Length*Width
	fmt.Println("The rectangular area is:", RectangularArea)
	

}