package main

import "fmt"
import "math"

type Circle struct {
	
	Radius float64
	CircleArea float64
	Circumference float64
}


func main() {
	
	var Radius float64 
	fmt.Println("Radius:")
	fmt.Scanln(&Radius)

	var  CircleArea float64
	var  Circumference float64

	CircleArea = math.Pi*Radius*Radius	
	Circumference = 2*math.Pi*Radius

	fmt.Println("The circle area is:", CircleArea )
	fmt.Println("The circle circumference area is:", Circumference)

}