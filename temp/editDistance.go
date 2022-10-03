package main

import (
	"fmt"
	"math"
)

func OneEditApart(s1, s2 string) bool {
	arrS1 := []byte(s1)
	arrS2 := []byte(s2)
	lenS1 := len(arrS1)
	lenS2 := len(arrS2)

	changeCount := 0
	exact := true

	if lenS2 == lenS1 {
		for i := 0; i < lenS1; i++ {
			if arrS1[i] != arrS2[i] {
				exact = false
				changeCount++
				if changeCount > 1 {
					return false
				}
			}
		}
	} else if math.Abs(float64(lenS2-lenS1)) == 1 {
		larger := 0
		if lenS1 > lenS2 {
			larger = lenS1
		} else {
			larger = lenS2
		}
		i, j := 0, 0
		for i < lenS1 && j < lenS2 {
			if arrS1[i] != arrS2[j] {
				exact = false
				if changeCount <= 1 {
					if larger == lenS1 {
						i++
					} else {
						j++
					}
				} else {
					return false
				}
				changeCount++
			}
			i++
			j++
		}
		if exact {
			return true
		}
	} else {
		return false
	}
	if !exact {
		return true
	} else {
		return false
	}

}

func main() {
	fmt.Println(OneEditApart("cat", "dog"))
	fmt.Println(OneEditApart("cat", "cats"))
	fmt.Println(OneEditApart("cat", "cut"))
	fmt.Println(OneEditApart("cat", "cast"))
	fmt.Println(OneEditApart("cat", "at"))
	fmt.Println(OneEditApart("cat", "act"))
}
