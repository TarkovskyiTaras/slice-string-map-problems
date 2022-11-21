package main

import "fmt"

/*
Given a number n, create and return a new int slice of length n, containing the numbers 0, 1, 2, ... n-1.
The given n may be 0, in which case just return a length 0 slice.

fizzSlice(4) → [0, 1, 2, 3]
fizzSlice(1) → [0]
fizzSlice(10) → [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
*/

func sliceMaker(numbOfElements int) []int {
	resultSlice := make([]int, numbOfElements, numbOfElements*2)
	for i := range resultSlice {
		resultSlice[i] = i
	}
	return resultSlice
}

func main() {
	fmt.Println(sliceMaker(5)) //[0 1 2 3 4]
}

//func fizzSlice(numbOfSliceEl int) {
//	resultSlice := make([]int, numbOfSliceEl)
//	elementValueSpacing := 0
//	for i, _ := range resultSlice {
//		resultSlice[i] = elementValueSpacing
//		elementValueSpacing++
//
//	}
//	fmt.Println(resultSlice)
//}
