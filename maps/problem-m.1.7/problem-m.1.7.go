package main

import "fmt"

/*
Modify and return the given map as follows: for this problem the map may or may not contain the "a" and "b" keys.
If both keys are present, append their 2 string values together and store the result under the key "ab".

mapAB({"a": "Hi", "b": "There"}) → {"a": "Hi", "ab": "HiThere", "b": "There"}
mapAB({"a": "Hi"}) → {"a": "Hi"}
mapAB({"b": "There"}) → {"b": "There"}
*/

func mapAB(m map[string]string) map[string]string {
	_, okA := m["a"]
	_, okB := m["b"]
	if okA && okB {
		m["ab"] = m["a"] + m["b"]
	}
	return m
}

func main() {
	fmt.Println(mapAB(map[string]string{"a": "aaaa", "b": "bbbb", "c": "ccccc", "d": "dddd"}))
	// map[a:aaaa ab:aaaabbbb b:bbbb c:ccccc d:dddd]
}
