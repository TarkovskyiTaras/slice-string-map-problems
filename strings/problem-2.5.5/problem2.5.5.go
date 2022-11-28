package main

import (
	"fmt"
	"log"
)

/*

Given a string, return a string length 1 from its front, unless front is false, in which case return a string length 1 from its back. The string will be non-empty.


theEnd("Hello", true) → "H"
theEnd("Hello", false) → "o"
theEnd("oh", true) → "o"
*/

func theEnd(givenString string, front bool) string {
	if len(givenString) < 1 {
		log.Fatal("invalid data entered")
	}
	if front {
		return string(givenString[0])
	} else {
		lastIndex := len(givenString) - 1
		return string(givenString[lastIndex])
	}
}

func main() {
	fmt.Println(theEnd("Taras", false))
}

//func theEnd(nonEmString string, boolFront bool) {
//	var resultString string
//	if boolFront == true {
//		resultString = string(nonEmString[0])
//	} else {
//		resultString = string(nonEmString[len(nonEmString)-1])
//	}
//	fmt.Println(resultString)
//}
