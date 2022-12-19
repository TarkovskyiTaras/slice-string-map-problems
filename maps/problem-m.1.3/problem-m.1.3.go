package main

import "fmt"

/*
Modify and return the given map as follows: if the keys "a" and "b" are both in the map and have equal values, remove them both.

mapAB2({"a": "aaa", "b": "aaa", "c": "cake"}) → {"c": "cake"}
mapAB2({"a": "aaa", "b": "bbb"}) → {"a": "aaa", "b": "bbb"}
mapAB2({"a": "aaa", "b": "bbb", "c": "aaa"}) → {"a": "aaa", "b": "bbb", "c": "aaa"}
*/

func mapAB2(m map[string]string) map[string]string {
	ma, ok1 := m["a"]
	mb, ok2 := m["b"]
	if ok1 && ok2 && ma == mb {
		delete(m, "a")
		delete(m, "b")
	}
	return m
}

func main() {
	fmt.Println(mapAB2(map[string]string{"a": "1", "c": "2"}))
	fmt.Println(mapAB2(map[string]string{"a": "1", "b": "1", "c": "3"}))
}
