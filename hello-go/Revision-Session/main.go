package main

import (
	"fmt"
	"introduction/add"
	"introduction/add/kl"
)

func main(){
	fmt.Println("Helo go!")
	add.Sum()
	kl.Kl()

    arr:=int{1,2,3,4,5,6}
	fmt.Println(arr)

	var slc [] int
	fmt.Println(slc == nil)

	slc= arr[3:5]
	fmt.Println(slc)

	slc1:= make([]int , 2, 4)
	fmt.Println(slc1)
	slc1= append(slc1,3)
	slc1= append(slc1,5)
	fmt.Println(slc1)
	fmt.Println(&slc1[0])
	mp:=make(map[string]string)
	mp["name"]="Neeraja"
	fmt.Println(mp)
	fmt.Println(mp[name])
}