package main

import "fmt"

/*
https://codingbat.com/prob/p127384

Return the sum of the numbers in the array, returning 0 for an empty array.
Except the number 13 is very unlucky, so it does not count and numbers
that come immediately after a 13 also do not count.

sum13([1, 2, 2, 1]) → 6
sum13([1, 1]) → 2
sum13([1, 2, 2, 1, 13]) → 6
*/

func sumOfSliceElementsNo13(givenSlice []int) int {
	sumOfElements := 0

	for i := range givenSlice {
		if givenSlice[i] == 13 {
			return sumOfElements
		}
		sumOfElements += givenSlice[i]
	}
	return sumOfElements
}

func main() {
	fmt.Println(sumOfSliceElementsNo13([]int{1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1})) //7
}

//func sumOfSliceElements(givenSlice []int) {
//	sumOfEl := 0
//	unluckyNumb := 13
//	for i, c := range givenSlice {
//		lastIndex := len(givenSlice) - 1
//		sumOfEl += c
//
//		if c == unluckyNumb && i == lastIndex {
//			sumOfEl -= givenSlice[lastIndex]
//
//		}
//
//		if c == unluckyNumb && i != lastIndex {
//			sumOfEl -= c
//
//			sumOfEl -= givenSlice[i+1]
//
//		}
//
//	}
//	fmt.Println(sumOfEl)
//}
