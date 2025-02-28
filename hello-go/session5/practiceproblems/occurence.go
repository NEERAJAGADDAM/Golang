package main

import (
	"fmt"
)

func RemoveOccurrences[T comparable](slice *[]T, value T) {
	newLength := 0
	for _, v := range *slice {
		if v != value {
			(*slice)[newLength] = v
			newLength++
		}
	}
	*slice = (*slice)[:newLength]
}

func main() {
	slice := []int{1, 2, 3, 4, 2, 5, 2, 6, 7}
	valueToRemove := 2
	RemoveOccurrences(&slice, valueToRemove)
	fmt.Println(slice) 
}
