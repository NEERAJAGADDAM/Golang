package main

import (
	"fmt"
	"strconv"
)

func digitFrequency(n int) map[rune]int {
	freqMap := make(map[rune]int)
	numStr := strconv.Itoa(abs(n)) 

	for _, digit := range numStr {
		freqMap[digit]++
	}
	return freqMap
}


func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	num := 1123581321
	freq := digitFrequency(num)

	for digit, count := range freq {
		fmt.Printf("%c: %d\n", digit, count)
	}
}
