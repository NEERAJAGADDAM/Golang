package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary int
}

func increaseSalary(emp *Employee, avg_Salary int, hike int) {
	if emp.Salary < avg_Salary {
		emp.Salary = emp.Salary + hike
		fmt.Println("Salary is updated for", emp.Name)
	} else {
		fmt.Println("There is no hike for", emp.Name)
	}
}

func main() {
	fmt.Println("Welcome to employee management")
	fmt.Println("***************************")

	var Id int
	fmt.Println("Please Enter your Id Number:")
	fmt.Scanln(&Id)

	var Name string
	fmt.Println("Please Enter your Name:")
	fmt.Scanln(&Name)

	var Salary int
	fmt.Println("Please Enter your Salary:")
	fmt.Scanln(&Salary)

	
	emp := Employee{Id: Id, Name: Name, Salary: Salary}

	
	avg_Salary := 40000
	hike := 5000

	
	increaseSalary(&emp, avg_Salary, hike)

	fmt.Println("\nUpdated Employee Details:")
	fmt.Println("Id:", emp.Id)
	fmt.Println("Name:", emp.Name)
	fmt.Println("Salary:", emp.Salary)
}

