package main

import "fmt"

func main() {

	str := "abcd"

	for i := 0; i <= len(str); i++ {
		for j := i + 1; j <= len(str); j++ {
			fmt.Println(str[i:j])
		}
	}

}
