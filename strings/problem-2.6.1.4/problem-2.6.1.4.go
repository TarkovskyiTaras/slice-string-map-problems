package main

import (
	"fmt"
	"log"
)

//Given a string and an int n, return a string made of n repetitions of the last n characters of the string.
//You may assume that n is between 0 and the length of the string, inclusive.

func endString(givenString string, repNumb uint) string {
	if int(repNumb) > len(givenString) {
		log.Fatal("invalid data passed")
	}
	var resultString string
	for i := 1; i <= int(repNumb); i++ {
		resultString += givenString[len(givenString)-int(repNumb):]
	}
	return resultString
}

func main() {
	fmt.Println(endString("taras", 4)) // "arasarasarasaras"
}

//func repeatEnd(givenString string, numbOfRep int) {
//
//	endString := givenString[len(givenString)-numbOfRep:]
//	var endStringNTimes string
//	for x := 1; x <= numbOfRep; x++ {
//		endStringNTimes += endString
//	}
//	fmt.Println(endStringNTimes)
//}
