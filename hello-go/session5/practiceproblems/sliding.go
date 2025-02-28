package main

import (
	"fmt"
)


func SlidingWindowMax(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}

	var result []int
	var deque []int 

	for i := 0; i < len(nums); i++ {
	
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

	
		deque = append(deque, i)

		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("Sliding Window Maximum:", SlidingWindowMax(nums, k))
}
