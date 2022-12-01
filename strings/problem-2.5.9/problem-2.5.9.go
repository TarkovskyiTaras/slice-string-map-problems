package main

import "fmt"

/*
Given a string, if the string begins with "red" or "blue" return that color string, otherwise return the empty string.

seeColor("redxx") → "red"
seeColor("xxred") → ""
seeColor("blueTimes") → "blue"
*/

func seeColor(givenString string) string {
	colors := []string{"red", "blue"}
	for _, cl := range colors {
		if givenString[:len(cl)] == cl {
			return cl
		}
	}
	return ""
}

func main() {
	fmt.Println(seeColor("redTaras")) // "red"
}

//func seeColor(givenString string) {
//	if strings.Contains(givenString[0:3], "red") {
//		fmt.Println("red")
//	} else if strings.Contains(givenString[0:4], "blue") {
//		fmt.Println("blue")
//	} else {
//		fmt.Println("")
//	}
//}
