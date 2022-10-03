// We don’t provide test cases in this language yet, but have outlined the signature for you. Please write your code below, and don’t forget to test edge cases!
package main

import "fmt"

func findEncryptedWord(s string) string {

	if len(s) <= 1 {
		return s
	}

	mid := 0
	if len(s)%2 == 0 {
		mid = (len(s) - 1) / 2
	} else {
		mid = (len(s) / 2)
	}
	fmt.Println((len(s) - 1) / 2)
	str := string(s[mid])
	str = str + findEncryptedWord(s[:mid])
	str = str + findEncryptedWord(s[mid+1:])
	return str
}

func main() {
	// Call findEncryptedWord() with test cases here
	s := "facebook"
	fmt.Println(findEncryptedWord(s))
}
