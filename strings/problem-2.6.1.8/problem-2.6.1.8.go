package main

import "fmt"

/*
Return the number of times that the string "hi" appears anywhere in the given string.

countHi("abc hi ho") → 1
countHi("ABChi hi") → 2
countHi("hihi") → 2
*/

func countHi(s, substring string) int {
	matchCount := 0
	t := len(s) - len(substring) + 1
	for i := 0; i < t; i++ {
		if s[i:i+len(substring)] == substring {
			matchCount++
		}
	}
	return matchCount
}

func main() {
	fmt.Println(countHi("tarasnikitatarasikitataras", "taras"))
}
