package main

import "fmt"

/*
Modify and return the given map as follows: if the key "a" has a value, set the key "b" to have that same value.
In all cases remove the key "c", leaving the rest of the map unchanged.

mapShare({"a": "aaa", "b": "bbb", "c": "ccc"}) → {"a": "aaa", "b": "aaa"}
mapShare({"b": "xyz", "c": "ccc"}) → {"b": "xyz"}
mapShare({"a": "aaa", "c": "meh", "d": "hi"}) → {"a": "aaa", "b": "aaa", "d": "hi"}
*/

func mapShare(m map[string]string) map[string]string {
	if _, ok := m["a"]; ok {
		m["b"] = m["a"]
	}
	delete(m, "c")
	return m
}

func main() {
	map1 := make(map[string]string)
	for x := 'a'; x <= 'z'; x++ {
		map1[string(x)] = string(x) + string(x) + string(x)
	}
	fmt.Print(mapShare(map1))
	// map[a:aaa b:aaa d:ddd e:eee f:fff g:ggg h:hhh i:iii j:jjj k:kkk l:lll m:mmm n:nnn o:ooo p:ppp q:qqq r:rrr s:sss t:ttt u:uuu v:vvv w:www x:xxx y:yyy z:zzz]
}
