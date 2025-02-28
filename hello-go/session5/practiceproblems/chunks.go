package main

import (
	"fmt"
)

// ChunkSlice splits a slice into smaller slices of size k
func ChunkSlice[T any](slice []T, k int) [][]T {
	if k <= 0 {
		return nil
	}

	var chunks [][]T
	for i := 0; i < len(slice); i += k {
		end := i + k
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 3
	result := ChunkSlice(slice, k)
	fmt.Println(result) // Output: [[1 2 3] [4 5 6] [7 8 9]]
}
