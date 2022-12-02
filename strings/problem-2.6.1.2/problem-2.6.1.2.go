package main

import "fmt"

/*
Return the number of times that the string "code" appears anywhere in the given string,
except we'll accept any letter for the 'd', so "cope" and "cooe" count.

countCode("aaacodebbb") → 1
countCode("codexxcode") → 2
countCode("cozexxcope") → 2
*/

func twoStringsEqual(s1, s2 string, anyCharIndex int) bool {
	for i := range s1 {
		if i != anyCharIndex && s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func countCode(givenString, subString string, anyCharIndex int) int {
	matchCounter := 0
	lenDiff := len(givenString) - len(subString)
	lenSubString := len(subString)
	for i := 0; i <= lenDiff; i++ {
		if twoStringsEqual(givenString[i:i+lenSubString], subString, anyCharIndex) {
			matchCounter++
		}
	}
	return matchCounter
}

func main() {
	fmt.Println(countCode("aaatarasaaatarasfasfaftarasdsafataraptasas", "tapas", 2))
}

//func main() {
//	givenString := "codeefeebcosebbsefsdcoee"
//	searchedString := "code"
//	var searchedStringDiffSecChar string
//	matchCount := 0
//	for i, _ := range givenString {
//		if i > len(givenString)-4 {
//			break
//		}
//		x := givenString[i : i+4]
//		fmt.Println("given string: ", x)
//		fmt.Println("searched string before if: ", searchedString)
//		if string(x[2]) != "d" {
//			searchedStringDiffSecChar = fmt.Sprintf("co%se", string(x[2]))
//		}
//		fmt.Println("searched string diff second char: ", searchedStringDiffSecChar)
//		if x == searchedString || x == searchedStringDiffSecChar {
//			matchCount++
//		}
//
//	}
//	fmt.Printf("Searched string has apeeared %d times in the given string.", matchCount)
//}
