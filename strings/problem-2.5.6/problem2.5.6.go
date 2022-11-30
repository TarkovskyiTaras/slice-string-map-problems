package main

import "fmt"

/*
Given a string, return true if it ends in "ly".

endsLy("oddly") → true
endsLy("y") → false
endsLy("oddy") → false
*/

func endsLy(givenString string) bool {
	return givenString[len(givenString)-2:] == "ly"
}

func main() {
	fmt.Println(endsLy("Tarasly"))
}

//func endsLy2(givenString string) {
//	if givenString[len(givenString)-2:len(givenString)] == "ly" {
//		fmt.Println(true)
//	} else {
//		fmt.Println(false)
//	}
//}
