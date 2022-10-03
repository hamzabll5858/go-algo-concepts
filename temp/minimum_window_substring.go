package main

import (
	"fmt"
	"math"
)

func SlidingWindow(s, t string) (int, int, int) {
	if t == "" {
		return 0, 0, 0
	}
	tLen := len(t)
	sLen := len(s)
	have, need := 0, len(t)
	start, end, lenghtTotal := -1, -1, math.MaxInt
	countT := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < tLen; i++ {
		countT[t[i]] = 1
	}

	j := 0
	for i := 0; i < sLen; i++ {
		c := s[i]
		window[c] = window[c] + 1
		if _, ok := countT[c]; ok {
			if window[c] == countT[c] {
				have++
			}
		}

		for have == need {
			if i-j+1 < lenghtTotal {
				start = j
				end = i
				lenghtTotal = i - j + 1
			}
			window[s[j]] = window[s[j]] - 1
			if _, ok := countT[s[j]]; ok {
				if window[s[j]] < countT[s[j]] {
					have--
				}
			}
			j++
		}

	}

	return start, end, lenghtTotal
}
func main() {

	s := "ADOBECODFBANC"
	t := "ABC"
	fmt.Println(SlidingWindow(s, t))
}
