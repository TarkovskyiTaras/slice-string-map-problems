package main

import "fmt"

/*
Given a string, return a string where for every char in the original, there are two chars.

doubleChar("The") → "TThhee"
doubleChar("AAbb") → "AAAAbbbb"
doubleChar("Hi-There") → "HHii--TThheerree"
*/

func doubleChar(givenString string) string {
	doubleCharString := ""
	for _, gs := range givenString {
		doubleCharString += string(gs) + string(gs)
	}
	return doubleCharString
}

func main() {
	fmt.Println(doubleChar("Taras"))
}

//func main() {
//	givenString := "tarasnikita"
//	var doubleCharString string
//	for i, _ := range givenString {
//		doubleCharString += string(givenString[i]) + string(givenString[i])
//
//	}
//	fmt.Println(doubleCharString)
//}
