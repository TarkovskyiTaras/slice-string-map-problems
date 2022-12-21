package main

import "fmt"

/*
Given a map of food keys and topping values, modify and return the map as follows: if the key "potato" has a value, set that as the
value for the key "fries". If the key "salad" has a value, set that as the value for the key "spinach".

topping3({"potato": "ketchup"}) → {"potato": "ketchup", "fries": "ketchup"}
topping3({"potato": "butter"}) → {"potato": "butter", "fries": "butter"}
topping3({"salad": "oil", "potato": "ketchup"}) → {"spinach": "oil", "salad": "oil", "potato": "ketchup", "fries": "ketchup"}
*/

func topping3(m map[string]string) map[string]string {
	if _, ok := m["potato"]; ok {
		m["fries"] = m["potato"]
	}
	if _, ok := m["salad"]; ok {
		m["spinach"] = m["salad"]
	}
	return m
}

func main() {
	fmt.Println(topping3(map[string]string{"1": "one", "salad": "two", "3": "three", "potato": "four", "5": "five"}))
	// map[1:one 3:three 5:five fries:four potato:four salad:two spinach:two]
}
