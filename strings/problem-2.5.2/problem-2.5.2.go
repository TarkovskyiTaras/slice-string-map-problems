package main

import (
	"fmt"
	"log"
)

/*
Given an "out" string of even length, such as "<<>>", and a word, return a new string where the word is in the middle of the out string,
e.g. "<<word>>".

makeOutWord("<<>>", "Yay") → "<<Yay>>"
makeOutWord("<<>>", "WooHoo") → "<<WooHoo>>"
makeOutWord("[[]]", "word") → "[[word]]"
*/

func makeOutWord(sidePart, middlePart string) string {
	if len(sidePart)%2 != 0 {
		log.Fatal("sidePart's length should be even")
	}
	return sidePart[:len(sidePart)/2] + middlePart + sidePart[len(sidePart)/2:]
}

func main() {
	fmt.Println(makeOutWord("<<||||>>", "Taras"))
}

//func makeOutWord(outString string, wordString string) {
//	resultString := outString[:2] + wordString + outString[2:]
//	fmt.Println(resultString)
//}
