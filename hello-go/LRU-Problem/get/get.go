package getting

import (
	"fmt"
	"LRU-Problem/put"
)

func Get(key int) int {
	if value, exists := adding.Add[key]; exists {
		fmt.Println("Fetched Value:", value)
		return value
	}
	fmt.Println("Key not found")
	return -1
}
