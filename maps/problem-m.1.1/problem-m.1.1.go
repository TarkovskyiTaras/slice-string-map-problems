package main

import "fmt"

/*
Modify and return the given map as follows: if the key "a" has a value, set the key "b" to have that value,
and set the key "a" to have the value "". Basically "b" is a bully, taking the value and replacing it with the empty string.

mapBully({"a": "candy", "b": "dirt"}) → {"a": "", "b": "candy"}
mapBully({"a": "candy"}) → {"a": "", "b": "candy"}
mapBully({"a": "candy", "b": "carrot", "c": "meh"}) → {"a": "", "b": "candy", "c": "meh"}
*/

func mapBully(m map[string]string) map[string]string {
	if m == nil {
		m = make(map[string]string)
	}
	if x, ok := m["a"]; ok {
		m["b"] = x
		m["a"] = ""
	}
	return m
}

func main() {
	fmt.Println(mapBully(map[string]string{"a": "taras", "b": "nikita"})) // map[a: b:taras]
	fmt.Println(mapBully(map[string]string{"b": "123", "c": "456"}))      // map[b:123 c:456]
}
