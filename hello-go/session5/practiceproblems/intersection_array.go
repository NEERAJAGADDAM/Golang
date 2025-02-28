package main

import (
	"fmt"
)


func intersection(arr1, arr2 []int) []int {
	freqMap := make(map[int]bool)
	result := []int{}


	for _, num := range arr1 {
		freqMap[num] = true
	}

	for _, num := range arr2 {
		if freqMap[num] {
			result = append(result, num)
			delete(freqMap, num) 
		}
	}

	return result
}

func main() {
	arr1 := []int{1, 2, 2, 3, 4, 5}
	arr2 := []int{2, 3, 5, 7}

	fmt.Println(intersection(arr1, arr2)) 
}
