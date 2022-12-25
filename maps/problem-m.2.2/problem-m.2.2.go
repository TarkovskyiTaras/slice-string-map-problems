package main

import "fmt"

/*
The classic word-count algorithm: given an array of strings, return a Map<String, Integer> with a key for each different
string, with the value the number of times that string appears in the array.

wordCount(["a", "b", "a", "c", "b"]) → {"a": 2, "b": 2, "c": 1}
wordCount(["c", "b", "a"]) → {"a": 1, "b": 1, "c": 1}
wordCount(["c", "c", "c", "c"]) → {"c": 4}
*/

func wordCount(ss []string) map[string]int {
	resultMap := make(map[string]int)
	for _, c := range ss {
		resultMap[c]++
	}
	return resultMap
}

func main() {
	fmt.Println(wordCount([]string{"a", "a", "b", "b", "c", "c", "c", "c", "d", "d", "e", "f", "f", "f", "g"}))
} // map[a:2 b:2 c:4 d:2 e:1 f:3 g:1]
