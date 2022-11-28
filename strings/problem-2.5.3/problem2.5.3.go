package main

import (
	"fmt"
	"log"
)

/*
Given a string of even length, return the first half. So the string "WooHoo" yields "Woo".

firstHalf("WooHoo") → "Woo"
firstHalf("HelloThere") → "Hello"
firstHalf("abcdef") → "abc"
*/

func firstHalf(givenString string) string {
	if len(givenString)%2 != 0 {
		log.Fatal("invalid data passed")
	}
	return givenString[:len(givenString)/2]
}

func main() {
	fmt.Println(firstHalf("Taras12345")) //Taras
}

//func firstHalf(stringOfEvenLength string) {
//	firstHalfString := stringOfEvenLength[:len(stringOfEvenLength)/2]
//	fmt.Println(firstHalfString)
//}
