package main

import "fmt"

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	mid := (len(arr)) / 2

	return Merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
func Merge(arr1 []int, arr2 []int) []int {
	totalLength := len(arr1) + len(arr2)
	result := make([]int, totalLength)
	left, right := 0, 0
	for i := 0; i < totalLength; i++ {
		if right >= len(arr2) && left < len(arr1) {
			result[i] = arr1[left]
			left++
		} else if left >= len(arr1) && right < len(arr2) {
			result[i] = arr2[right]
			right++
		} else if arr1[left] < arr2[right] {
			result[i] = arr1[left]
			left++
		} else {
			result[i] = arr2[right]
			right++
		}
	}
	return result
}
func main() {
	temp := []int{1, 5, 3, 2, 5, 22, 6, 8, 2, 22}
	temp = MergeSort(temp)
	fmt.Println(temp)
}
