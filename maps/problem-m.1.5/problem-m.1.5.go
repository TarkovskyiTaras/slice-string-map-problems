package main

import "fmt"

/*
Given a map of food keys and their topping values, modify and return the map as follows: if the key "ice cream" has a value,
set that as the value for the key "yogurt" also. If the key "spinach" has a value, change that value to "nuts".

topping2({"ice cream": "cherry"}) → {"yogurt": "cherry", "ice cream": "cherry"}
topping2({"spinach": "dirt", "ice cream": "cherry"}) → {"yogurt": "cherry", "spinach": "nuts", "ice cream": "cherry"}
topping2({"yogurt": "salt"}) → {"yogurt": "salt"}
*/

func topping2(m map[string]string) map[string]string {
	if _, ok := m["ice cream"]; ok {
		m["yogurt"] = m["ice cream"]
	}
	if _, ok := m["spinach"]; ok {
		m["spinach"] = "nuts"
	}
	return m
}

func main() {
	fmt.Print(topping2(map[string]string{"ice cream": "value for ice cream", "yogurt": "value for yogurt",
		"spinach": "value for spinach"})) // map[ice cream:value for ice cream spinach:nuts yogurt:value for ice cream]
}
