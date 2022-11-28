package main

import (
	"fmt"
	"log"
)

/*
Given 2 strings, return their concatenation, except omit the first char of each. The strings will be at least length 1.

nonStart("Hello", "There") → "ellohere"
nonStart("java", "code") → "avaode"
nonStart("shotl", "java") → "hotlava"
*/

func nonStart(firstString, secondString string) string {
	if len(firstString) < 1 || len(secondString) < 1 {
		log.Fatal("string length less than 1")
	}
	return firstString[1:] + secondString[1:]
}

func main() {
	fmt.Println(nonStart("taras", "cool"))
}

//func nonStart(firstString string, secondString string) {
//	resultString := firstString[1:] + secondString[1:]
//	fmt.Println(resultString)
//}
