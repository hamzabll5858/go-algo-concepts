package main

import (
	"fmt"
	"sort"
)

// Write any import statements here

func pairSums(k int32, n int32, arr []int32) int {

	compliment := make(map[int32]int32)
	indexes := make(map[int32]int32)

	for i := int32(0); i < n; i++ {
		compliment[k-arr[i]] = i
	}

	for i := int32(0); i < n; i++ {
		if _, ok := compliment[arr[i]]; ok {
			if compliment[arr[i]] != i {
				indexes[compliment[arr[i]]] = i
			}
		}
	}

	for key, value := range indexes {

		if _, ok := indexes[value]; ok && indexes[value] == key {
			delete(indexes, key)
		} else {
			fmt.Println(key, value)
		}

	}

	// Write your code here
	return len(indexes)
}

func main() {
	arr := []int{1, 2, 3, 4, 3}

	sort.Ints(arr)
	fmt.Println(arr)
	//pairSums(k, n, arr)
}
