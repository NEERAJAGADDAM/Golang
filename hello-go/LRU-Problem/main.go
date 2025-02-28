package main

import (
	"fmt"
	"LRU-Problem/put"
	"LRU-Problem/get"
)

func main() {
	adding.Put(1,1)
	adding.Put(2,2)

	fmt.Println("Getting Value for key 3:", getting.Get(3))

	adding.Put(3,3)
	
	fmt.Println("Getting Value for key 3:", getting.Get(3))
	
	adding.Put(4,4)
	adding.Put(5,5)
	adding.Put(6,6) 

	fmt.Println("Getting Value for key 3:", getting.Get(3))
}
