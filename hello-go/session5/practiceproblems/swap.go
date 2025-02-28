package main

import "fmt"

func main(){

	var a = map[string]string{"brand":"Motorola", "model":"C33"}
    var b = map[string]string{"brand":"Realme", "model": "C55"}

	fmt.Println("Before Swap:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)

     c:=a
	 a=b
	 b=c

	 fmt.Println("\nAfter Swap:")
	 fmt.Println("a:", a)
	 fmt.Println("b:", b)
}