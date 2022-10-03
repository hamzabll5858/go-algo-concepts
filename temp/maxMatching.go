// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func checkSameIndex(hash map[string]string) bool {
	for key, _ := range hash {
		if _, ok := hash[hash[key]]; ok {
			if key == hash[hash[key]] {
				return true
			}
		}
	}
	return false
}

func matchingPairs(s string, t string) int {

	matches := 0
	hash := make(map[string]string)
	duplicates := make(map[string]int)
	isdistinct := true

	for i := 0; i < len(s); i++ {
		if s[i] == t[i] {
			matches++
		} else {
			hash[string(s[i])] = string(t[i])
		}
		duplicates[string(t[i])]++
		if duplicates[string(t[i])] > 1 {
			isdistinct = false
		}
	}

	if matches == len(s) {
		if isdistinct {
			return matches - 2
		} else {
			return matches
		}
	} else {
		if checkSameIndex(hash) {
			if isdistinct {
				return matches + 2
			} else {
				return matches
			}
		} else {
			if isdistinct {
				return matches - 2
			} else {
				return matches
			}
		}
	}

}

func main() {
	// Call matchingPairs() with test cases here
	s := "abcda"
	t := "aecda"
	fmt.Println(matchingPairs(s, t))
}
