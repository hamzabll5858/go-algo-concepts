package main

import "fmt"

func BubbleSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		swapped := false
		for j := 0; j < arrLen-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
	}
	return arr
}

func SelectionSort(arr []int) []int {
	arrLen := len(arr)
	for i := 0; i < arrLen; i++ {
		for j := i; j < arrLen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func BinarySearch(arr []int, start int, end int, element int) bool {
	if start <= end {
		mid := (start + end) / 2
		if arr[mid] == element {
			return true
		} else if element < arr[mid] {
			return BinarySearch(arr, 1, mid-1, element)
		} else {
			return BinarySearch(arr, mid+1, end, element)
		}
	} else {
		return false
	}

}

func main() {
	arr := []int{2, 3, 5, 1, 55, 32, 11, 23, 0, 1}
	arr = SelectionSort(arr)
	fmt.Println(BinarySearch(arr, 0, len(arr)-1, 11))
	fmt.Println(arr)
}
