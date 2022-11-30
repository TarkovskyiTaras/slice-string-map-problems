package main

import (
	"fmt"
	"log"
)

/*
Given a string of odd length, return the string length 3 from its middle, so "Candy" yields "and".
The string length will be at least 3.

middleThree("Candy") → "and"
middleThree("and") → "and"
middleThree("solving") → "lvi"
*/

func middleThree(givenString string) string {
	if len(givenString)%2 != 1 {
		log.Fatal("string is supposed to be of odd length")
	}
	return givenString[len(givenString)/2-1 : len(givenString)/2+2]
}

func main() {
	fmt.Println(middleThree("Taras"))
}

//func middleThree(givenString string) {
//	if len(givenString) > 3 {
//		resultString := givenString[len(givenString)/2-1 : len(givenString)/2+2]
//		fmt.Println(resultString)
//	}
//}
