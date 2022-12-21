package main

import "fmt"

/*
Modify and return the given map as follows: if the keys "a" and "b" have values that have different lengths, then set "c"
to have the longer value. If the values exist and have the same length, change them both to the empty string in the map.

mapAB4({"a": "aaa", "b": "bb", "c": "cake"}) → {"a": "aaa", "b": "bb", "c": "aaa"}
mapAB4({"a": "aa", "b": "bbb", "c": "cake"}) → {"a": "aa", "b": "bbb", "c": "bbb"}
mapAB4({"a": "aa", "b": "bbb"}) → {"a": "aa", "b": "bbb", "c": "bbb"}
*/

func mapAB4(m map[string]string) map[string]string {
	_, ok1 := m["a"]
	_, ok2 := m["b"]

	if ok1 && ok2 && len(m["a"]) > len(m["b"]) {
		m["c"] = m["a"]
	}

	if ok1 && ok2 && len(m["b"]) > len(m["a"]) {
		m["c"] = m["b"]
	}

	if ok1 && ok2 && len(m["a"]) == len(m["b"]) {
		m["a"], m["b"] = "", ""
	}
	return m
}

func main() {
	fmt.Println(mapAB4(map[string]string{"a": "aaa", "b": "bbbb", "c": "ccc", "d": "ddd", "e": "eee"}))
	// map[a:aaa b:bbbb c:bbbb d:ddd e:eee]
}
