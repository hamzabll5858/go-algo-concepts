package main

import (
	"fmt"
	"math/rand"
)

func quicksort(arr []int) []int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr
	}
	left, right := 0, arrLen-1
	pivot := rand.Int()%arrLen - 1

	arr[pivot], arr[right] = arr[right], arr[pivot]
	for i := 0; i < arrLen; i++ {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[left], arr[right] = arr[right], arr[left]
	quicksort(arr[:left])
	quicksort(arr[left+1:])
	return arr
}

func main() {

	temp := []int{1, 35, 4, 2, 4, 66, 32, 10, 2}
	temp = quicksort(temp)
	fmt.Println(temp)
}
