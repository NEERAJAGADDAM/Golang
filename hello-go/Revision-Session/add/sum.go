package add

import "fmt"

func Sum(){
	var a,b int
	fmt.Println("Enter  a:")
	fmt.Scanln(&a)

	fmt.Println("Enter  b:")
	fmt.Scanln(&b)

	sum:= a+b
	fmt.Println(sum)  
}