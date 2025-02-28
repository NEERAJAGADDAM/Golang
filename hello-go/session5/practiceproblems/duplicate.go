package main

import "fmt"

func findDuplicates(nums []int) []int {
	freqMap := make(map[int]int)
	duplicates := []int{}

	for _, num := range nums {
		freqMap[num]++
	}

	for num, count := range freqMap {
		if count > 1 {
			duplicates = append(duplicates, num)
		}
	}

	return duplicates
}


func main() {
	nums := []int{1, 2, 3, 2, 4, 5, 6, 4, 4, 7}
	fmt.Println(findDuplicates(nums)) 
}