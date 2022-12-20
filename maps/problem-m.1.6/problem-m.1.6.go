package main

import "fmt"

/*
Modify and return the given map as follows: if exactly one of the keys "a" or "b" has a value in the map (but not both),
set the other to have that same value in the map.

mapAB3({"a": "aaa", "c": "cake"}) → {"a": "aaa", "b": "aaa", "c": "cake"}
mapAB3({"b": "bbb", "c": "cake"}) → {"a": "bbb", "b": "bbb", "c": "cake"}
mapAB3({"a": "aaa", "b": "bbb", "c": "cake"}) → {"a": "aaa", "b": "bbb", "c": "cake"}
*/

func mapAB3(m map[string]string) map[string]string {
	_, okA := m["a"]
	_, okB := m["b"]
	if okA && !okB {
		m["b"] = m["a"]
	}
	if okB && !okA {
		m["a"] = m["b"]
	}
	return m
}

func main() {
	map1 := map[string]string{"a": "aaa", "c": "cccc"}
	fmt.Println(mapAB3(map1)) //map[a:aaa b:aaa c:cccc]
}
