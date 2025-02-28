package main

import (
	"fmt"
	"LRU-Problem/put"
	"LRU-Problem/get"
)

func main() {
	put.Put(1, 10)
	put.Put(2, 20)
	put.Put(3, 30)
	put.Put(4, 40)

	fmt.Println("Getting Value for key 2:", get.Get(2))
	fmt.Println("Getting Value for key 1:", get.Get(1))
	fmt.Println("Getting Value for key 3:", get.Get(3))
}
