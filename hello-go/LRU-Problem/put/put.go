package put

import (
	"fmt"
	"LRU-Problem/remove"
)

const capacity = 3

var Add = make(map[int]int)
var Order []int

func Put(key int, value int) {
	if _, exists := Add[key]; exists {
		Add[key] = value
		Order = remove.RemoveKeyFromSlice(Order, key)
		Order = append(Order, key)
		return
	}

	if len(Add) >= capacity {
		lruKey := Order[0]
		delete(Add, lruKey)
		Order = Order[1:]
	}

	Add[key] = value
	Order = append(Order, key)

	fmt.Println("Updated Map:", Add)
	fmt.Println("Keys Order:", Order)
}

