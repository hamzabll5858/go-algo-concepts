package main

import "fmt"

func InsertionSort(arr []int) []int {

	j := 0
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		j = i
		for j > 0 && arr[j-1] > temp {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
	}
	return arr
}
func main() {
	temp := []int{1, 5, 3, 2, 5, 22, 6, 8, 2, 22}
	temp = InsertionSort(temp)
	fmt.Println(temp)
}
