package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'palindromeIndex' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func palindromeIndex(s string) int32 {
	// Write your code here

	i, j := int32(0), int32(len(s)-1)
	isPalindrom := true
	index := int32(-1)

	for i <= j && i <= int32(len(s)-1) && j <= int32(len(s)-1) {
		leftChar, RightChar := s[i], s[j]
		fmt.Println(string(RightChar), ",", string(leftChar))
		if s[i] == s[j] {
			i++
			j--
		} else if s[i] != s[j] {
			isPalindrom = false
			if s[i+1] != s[j] && s[j-1] != s[i] {
				return -1
			} else {
				if s[i+1] == s[j] {
					index = i
					i++
				} else if s[j-1] == s[i] {
					index = j
					j--
				}
			}
		}
	}

	if isPalindrom {
		return -1
	} else {
		return index
	}

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("C:\\Users\\hamza\\Desktop\\GoLang\\boilerplate\\out.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := palindromeIndex(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
