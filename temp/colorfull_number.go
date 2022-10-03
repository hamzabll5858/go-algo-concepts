package main

import "fmt"

func isColorful(num int) bool {
	digits := []int{}
	hashTable := make(map[int]int)
	powerSet := [][]int{{}}
	for num >= 1 {
		digits = append(digits, num%10)
		powerSet = append(powerSet, []int{num % 10})
		num = num / 10
	}

	for i := 0; i < len(digits); i++ {
		for j := len(digits); j > i+1; j-- {
			powerSet = append(powerSet, digits[i:j])
		}
	}

	for i := 0; i < len(powerSet); i++ {
		if len(powerSet[i]) != len(digits) || len(powerSet[i]) != 0 {
			tempMul := 1
			for j := 0; j < len(powerSet[i]); j++ {
				tempMul = tempMul * powerSet[i][j]
			}
			if _, ok := hashTable[tempMul]; ok {
				return false
			} else {
				hashTable[tempMul] = 0
			}
		}
	}

	fmt.Println(powerSet)
	return true
}
func main() {
	num := 3245
	fmt.Println(isColorful(num))

}
