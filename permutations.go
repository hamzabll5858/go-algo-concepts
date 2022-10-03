// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func minOperations(arr []int) int {

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j <= len(arr); j++ {
			fmt.Println(arr[i:j])
		}
	}
	// Write your code here
	return 0
}

func main() {
	arr := []int{3, 1, 2}
	fmt.Println(arr[0:3])
	minOperations(arr)
	// Call minOperations() with test cases here
}
