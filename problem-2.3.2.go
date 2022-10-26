package main

import (
	"fmt"
	"log"
)

/*
Given 2 arrays of ints, a and b, return true if they have the same first element or they have the same last element. Both arrays will be length 1 or more.
commonEnd([1, 2, 3], [7, 3]) → true
commonEnd([1, 2, 3], [7, 3, 2]) → false
commonEnd([1, 2, 3], [1, 3]) → true
*/

func twoSlicesOfSameLengthChecker(sliceOne []int, sliceTwo []int) bool {
	if len(sliceOne) == len(sliceTwo) {
		return true
	} else {
		return false
	}
}

func firstOrLastElementChecker(sliceOne []int, sliceTwo []int) bool {
	lastIndex := len(sliceOne) - 1
	if sliceOne[0] == sliceTwo[0] || sliceOne[lastIndex] == sliceTwo[lastIndex] {
		return true
	} else {
		return false
	}
}

func commonFirstOrLastElement(sliceOne []int, sliceTwo []int) bool {
	if twoSlicesOfSameLengthChecker(sliceOne, sliceTwo) == false {
		log.Fatal("entered slices in the arguments have different length")
	}
	return firstOrLastElementChecker(sliceOne, sliceTwo)
}

func main() {
	fmt.Println(commonFirstOrLastElement([]int{1, 3, 4}, []int{1, 3, 4}))
}

//package main
//
//import "fmt"
//
//func commonEnd(x1 int, y1 int, x2 int, y2 int) bool {
//	z := false
//	array1 := [4]int{x1, 2, 5, y1}
//	array2 := [4]int{x2, 41, 56, y2}
//
//	if array1[0] == array2[0] || array1[3] == array2[3] {
//		z = true
//	}
//	return z
//}
//
//func main() {
//	fmt.Println(commonEnd(22, 42, 24, 42))
//
//}
