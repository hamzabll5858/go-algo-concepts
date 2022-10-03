package main

import (
	"fmt"
	"strconv"
)

func spiral(n int) {

	matrix := make([][]string, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]string, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrix[i][j] = "0"
		}
	}

	maxIterate := n * n
	arrRow := []int{0, 1, 0, -1}
	arrCol := []int{1, 0, -1, 0}
	position := 0
	i, j := 0, 0
	count := 0
	for count < maxIterate {
		count++
		if matrix[i][j] == "0" {
			matrix[i][j] = strconv.Itoa(count)
		}

		if (i+arrRow[position] > n-1) || (j+arrCol[position] > n-1) || (i+arrRow[position] < 0) || (j+arrCol[position] < 0) || matrix[i+arrRow[position]][j+arrCol[position]] != "0" {
			position = (position + 1) % 4
		}
		i = i + arrRow[position]
		j = j + arrCol[position]
	}

	fmt.Println(matrix)
	for i := 0; i < n; i++ {
		fmt.Println(matrix[i])
	}
}

func main() {
	spiral(5)
}
