package main

import "fmt"

type Student struct {
	Name  string
	Marks float64
}

func main() {
	fmt.Println("Welcome to Student Management")
	fmt.Println("***************************")

	var Name string
	fmt.Println("Please Enter your Name:")
	fmt.Scanln(&Name)

	var Marks float64 
	fmt.Println("Please Enter your Marks:")
	fmt.Scanln(&Marks)

	
	if Marks > 90 {
		fmt.Println("Grade: A+")
	} else if Marks > 80 && Marks <= 90 {
		fmt.Println("Grade: A")
	} else if Marks > 70 && Marks <= 80 {
		fmt.Println("Grade: B+")
	} else if Marks > 60 && Marks <= 70 {
		fmt.Println("Grade: B")
	} else if Marks > 50 && Marks <= 60 {
		fmt.Println("Grade: C")
	} else if Marks > 40 && Marks <= 50 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: F")
	}
}
