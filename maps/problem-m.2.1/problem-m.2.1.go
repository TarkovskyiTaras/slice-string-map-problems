package main

import "fmt"

/*
Given an array of strings, return a Map<String, Integer> containing a key for every different string in the array,
always with the value 0. For example the string "hello" makes the pair "hello":0. We'll do more complicated counting later,
but for this problem the value is simply 0.

word0(["a", "b", "a", "b"]) → {"a": 0, "b": 0}
word0(["a", "b", "a", "c", "b"]) → {"a": 0, "b": 0, "c": 0}
word0(["c", "b", "a"]) → {"a": 0, "b": 0, "c": 0}
*/

func word0(ss []string) map[string]int {
	resMap := make(map[string]int)
	for _, c := range ss {
		resMap[c] = 0
	}
	return resMap
}

func main() {
	fmt.Println(word0([]string{"taras1", "taras2", "taras3", "taras4", "taras5"}))
}
