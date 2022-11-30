package main

import "fmt"

/*
Given 2 strings, a and b, return a new string made of the first char of a and the last char of
b, so "yo" and "java" yields "ya". If either string is length 0, use '@' for its missing char.

lastChars("last", "chars") → "ls"
lastChars("yo", "java") → "ya"
lastChars("hi", "") → "h@"
*/

func lastChars(stringA, stringB string) string {
	if len(stringA) == 0 {
		return "@" + string(stringB[len(stringB)-1])
	}
	if len(stringB) == 0 {
		return string(stringA[0]) + "@"
	}
	return string(stringA[0]) + string(stringB[len(stringB)-1])
}

func main() {
	fmt.Println(lastChars("taras", "cool"))
}

//func lastChars(stringA string, stringB string) {
//	resultString := string(stringA[0]) + string(stringB[len(stringB)-1])
//	if len(stringA) == 0 {
//		stringA = "@"
//	}
//	if len(stringB) == 0 {
//		stringB = "@"
//	}
//	fmt.Println(resultString)
//}
