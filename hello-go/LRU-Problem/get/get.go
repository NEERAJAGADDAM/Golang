package get

import (
	"fmt"
	"LRU-Problem/put"
	"LRU-Problem/remove"
)

func Get(key int) int {
	if value, exists := put.Add[key]; exists {
		put.Order = remove.RemoveKeyFromSlice(put.Order, key)
		put.Order = append(put.Order, key)

		fmt.Println("Fetched Value:", value)
		fmt.Println("Updated Keys Order:", put.Order)
		return value
	}
	fmt.Println("Key not found")
	return -1
}
