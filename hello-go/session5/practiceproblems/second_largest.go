package main

import "fmt"

func main(){
	nums:= []int{1,2,4,6,8,3,9}
    var c =0

	for i:=0;i<len(nums);i++{
		if nums[i]>c{
             c=nums[i] 
		}
	}

	fmt.Println(c)
}