// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func areTheyEqual(arr_1 []int, arr_2 []int) bool {
	// Write your code here
	length := len(arr_1)
	startIndex := 0
	endIndex := length - 1

	reverseStrStart := 0
	reverseStrEnd := 0

	startSet := false
	endSet := false

	for startIndex <= endIndex && !startSet || !endSet {
		if arr_1[startIndex] == arr_2[startIndex] {
			startIndex++
		} else {
			reverseStrStart = startIndex
			startSet = true
		}

		if arr_1[endIndex] == arr_2[endIndex] {
			endIndex--
		} else {
			reverseStrEnd = endIndex
			endSet = true
		}
	}

	arr_2 = strReverse(arr_2, reverseStrStart, reverseStrEnd)

	if checkEquals(arr_1, arr_2) {
		return true
	} else {
		return false
	}
}

func strReverse(arr []int, start, end int) []int {
	i := start
	j := end
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
}

func checkEquals(arr_1 []int, arr_2 []int) bool {
	for i := 0; i < len(arr_1); i++ {
		if arr_1[i] != arr_2[i] {
			return false
		}
	}
	return true
}

func main() {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	B := []int{1, 2, 7, 6, 5, 4, 3, 8, 9, 10}
	fmt.Println(areTheyEqual(A, B))
	// Call areTheyEqual() with test cases here
}
