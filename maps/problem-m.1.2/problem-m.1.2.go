package main

import "fmt"

/*
Given a map of food keys and topping values, modify and return the map as follows: if the key "ice cream" is present,
set its value to "cherry". In all cases, set the key "bread" to have the value "butter".

topping1({"ice cream": "peanuts"}) → {"bread": "butter", "ice cream": "cherry"}
topping1({}) → {"bread": "butter"}
topping1({"pancake": "syrup"}) → {"bread": "butter", "pancake": "syrup"}
*/

func topping1(m map[string]string) map[string]string {
	if _, ok := m["ice cream"]; ok {
		m["ice cream"] = "cherry"
	}
	m["bread"] = "butter"
	return m
}

func main() {
	fmt.Println(topping1(map[string]string{"ice cream": "vanilla", "pancake": "syrup"}))
}
