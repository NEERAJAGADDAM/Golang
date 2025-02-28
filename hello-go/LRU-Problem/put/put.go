package adding

import "fmt"

var Add = make(map[int]int)

func Put(key int, value int) {
	const capacity = 3
	if len(Add) >=capacity {
		for k := range Add {
			delete(Add, k) 
			break
		}
	}
	Add[key] = value
	fmt.Println("Updated Map:", Add)
}


