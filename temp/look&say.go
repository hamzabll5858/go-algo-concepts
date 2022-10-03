package main

import (
	"fmt"
	"strconv"
)

func lookSay(n int) {
	init := "1"
	currStr := []byte(init)
	for i := 0; i < n; i++ {
		var newStr string
		countCurr := 0
		prev := currStr[0]
		for _, value := range currStr {
			if prev != value {
				str := strconv.Itoa(countCurr) + string(prev)
				newStr = newStr + str
				countCurr = 0
			}
			prev = value
			countCurr++
		}
		str := strconv.Itoa(countCurr) + string(prev)
		newStr = newStr + str
		currStr = []byte(newStr)
		fmt.Println(string(newStr))

	}
}

func main() {
	lookSay(5)
}
