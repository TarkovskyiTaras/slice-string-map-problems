package main

import (
	"fmt"
	"log"
)

/*
Given a slice of ints of even length, return a new slice length 2 containing the middle two elements from the original
slice. The original slice will be length 2 or more.

makeMiddle([1, 2, 3, 4]) → [2, 3]
makeMiddle([7, 1, 2, 3, 4, 9]) → [2, 3]
makeMiddle([1, 2]) → [1, 2]
*/

func evenLengthOK(givenSlice []int) bool {
	return len(givenSlice)%2 == 0
}
func lengthTwoOrMoreOK(givenSlice []int) bool {
	return len(givenSlice) >= 2
}
func makeMiddle(givenSlice []int) []int {
	if !evenLengthOK(givenSlice) || !lengthTwoOrMoreOK(givenSlice) {
		log.Fatal("invalid data has been passed")
	}
	middleIndex := len(givenSlice) / 2
	resultSlice := make([]int, 2)[:2]
	resultSlice[0], resultSlice[1] = givenSlice[middleIndex-1], givenSlice[middleIndex]
	return resultSlice
}

func main() {
	fmt.Println(makeMiddle([]int{1, 2, 3, 4, 5, 6})) //[3, 4]
}

//func main() {
//	array1 := [8]int{2, 3, 6, 77, 89, 24, 35, 23}
//	x := len(array1) / 2
//	array2 := [2]int{array1[x-1], array1[x]}
//	fmt.Println(array1)
//	fmt.Println(array2)
//}
