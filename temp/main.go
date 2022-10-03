package main

import (
	"fmt"
	"regexp"
)

func main() {

	/*	g := graph.FillGraph()
		g.PrintAdjList()

		g.DFS()*/

	//fmt.Print(regexpMatch("(ABCD){3}.*\\-", "ABCDABCDABCDhamza-"))

	//fmt.Println(dynamic_programming.GetFibonacciNum(50, make(map[int]int)))
	/*	var t tree.Tree
		t.Insert(2)
		t.Insert(4)
		t.Insert(1)
		t.Insert(6)
		t.Insert(11)
		t.Insert(3)

		t.PostOrderParse(t.GetRoot())*/

}

func regexpMatch(pattern string, value string) bool {
	r, _ := regexp.Compile(pattern)

	r2 := regexp.MustCompile(`([0-9]{2}):([0-9]{2}):([0-9]{2})(AM|PM)`)
	matches := r2.FindStringSubmatch("12:40:22AM")

	hours := matches[1]
	minutes := matches[2]
	seconds := matches[3]
	fmt.Printf("%s:%s:%s", hours, minutes, seconds)

	return r.MatchString(value)
}
