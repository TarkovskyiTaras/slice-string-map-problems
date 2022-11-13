package main

import (
	"fmt"
)

/*
Given 2 slices of ints, a and b, return true if they have the same first element, or they have the same last element. Both slices will be length 1 or more.
commonEnd([1, 2, 3], [7, 3]) → true
commonEnd([1, 2, 3], [7, 3, 2]) → false
commonEnd([1, 2, 3], [1, 3]) → true
*/

func commonFirstOrLastElement(sliceOne []int, sliceTwo []int) bool {
	lastIndexSliceOne := len(sliceOne) - 1
	lastIndexSliceTwo := len(sliceTwo) - 1
	if sliceOne[0] == sliceTwo[0] || sliceOne[lastIndexSliceOne] == sliceTwo[lastIndexSliceTwo] {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(commonFirstOrLastElement([]int{1, 3, 4}, []int{2, 44, 3, -1, 4})) //true
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
