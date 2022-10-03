// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func checkSum(hash map[float64]float64) float64 {
	result := float64(0)
	for _, value := range hash {
		result = result + value
	}
	return result
}

func getBillionUsersDay(arr []float64) int {
	// Write your code here

	day := 0
	isBillion := false

	hash := make(map[float64]float64)

	for !isBillion {
		day++
		for i := 0; i < len(arr); i++ {
			if _, ok := hash[arr[i]]; ok {
				hash[arr[i]] = hash[arr[i]] * arr[i]
			} else {
				hash[arr[i]] = arr[i]
			}

		}
		if checkSum(hash) > float64(1000000000) {
			isBillion = true
		}
		fmt.Println(day)

	}

	return day
}

func main() {

	growthRates := []float64{1.01, 1.02}
	getBillionUsersDay(growthRates)
	// Call getBillionUsersDay() with test cases here
}
