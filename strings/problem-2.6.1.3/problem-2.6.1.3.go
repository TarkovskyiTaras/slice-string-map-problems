package main

import "fmt"

/*
Return true if the given string contains a "bob" string, but where the middle 'o' char can be any char.

bobThere("abcbob") → true
bobThere("b9b") → true
bobThere("bac") → false
*/

func subStringThere(givenString, subString string, anyCharIndex int) int {
	matchCount := 0
	for i := 0; i <= len(givenString)-len(subString); i++ {
		gsModified := givenString[i:i+anyCharIndex] + givenString[i+anyCharIndex+1:i+len(subString)]
		ssModified := subString[:anyCharIndex] + subString[anyCharIndex+1:]
		if gsModified == ssModified {
			matchCount++
		}
	}
	return matchCount
}

func main() {
	fmt.Println(subStringThere("bob123barfffbibfffbabwww", "bob", 1)) //3
}

//func main() {
//	givenString := "codeefeebbsebcbsefsdcoee"
//	searchedString := "bob"
//	matchStatus := false
//	for i, _ := range givenString {
//		if i > len(givenString)-3 {
//			break
//		}
//		x := givenString[i : i+3]
//		if x[0] == searchedString[0] && x[2] == searchedString[2] {
//			matchStatus = true
//			break
//		}
//
//	}
//	fmt.Println(matchStatus)
//}
