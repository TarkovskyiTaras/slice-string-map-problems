package main

import (
	"fmt"
	"log"
	"strings"
)

/*
Given a string, consider the prefix string made of the first N chars of the string.
Does that prefix string appear somewhere else in the string?
Assume that the string is not empty and that N is in the range 1..str.length().

prefixAgain("abXYabc", 1) → true
prefixAgain("abXYabc", 2) → true
prefixAgain("abXYabc", 3) → false
*/

func prefixAgain(givenString string, prefLength int) bool {
	if len(givenString) < 1 || prefLength < 1 {
		log.Fatal("invalid data has been entered.")
	}
	if strings.Contains(givenString[prefLength-1:], givenString[:prefLength]) {
		return true
	}
	return false
}

func main() {
	fmt.Println(prefixAgain("tratra12tra3", 3))
}

//func prefixAgain(givenString string, numbOfChars int) {
//	prefixString := givenString[:numbOfChars]
//	result := strings.Contains(givenString[numbOfChars:], prefixString)
//	fmt.Println(result)
//}
