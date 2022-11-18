package main

import (
	"fmt"
	"log"
)

/*
Given a slice of ints of odd length, return a new slice length "x" containing the elements from the middle of the slice.
The slice length will be at least 3.
*/

func oddLengthOK(givenSlice []int, middleSpan int) bool {
	return len(givenSlice)%2 != 0 && middleSpan%2 != 0
}

func midSpanNotMoreThanSliceOK(givenSlice []int, middleSpan int) bool {
	return len(givenSlice) >= middleSpan
}
func sliceLenNotLessThanThreeOK(givenSlice []int, middleSpan int) bool {
	return len(givenSlice) >= 3
}

func middleSlice(givenSlice []int, middleSpan int) []int {
	if !oddLengthOK(givenSlice, middleSpan) ||
		!midSpanNotMoreThanSliceOK(givenSlice, middleSpan) ||
		!sliceLenNotLessThanThreeOK(givenSlice, middleSpan) {
		log.Fatal("incorrect data has been passed")
	}
	resultSlice := make([]int, middleSpan)[:middleSpan]
	startPoint := (len(givenSlice) - middleSpan) / 2
	endPoint := startPoint + middleSpan

	for i, j := startPoint, 0; i < endPoint; i, j = i+1, j+1 {
		resultSlice[j] = givenSlice[i]
	}

	return resultSlice
}

func main() {
	fmt.Println(middleSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 15))
}

//func main() {
//	array1 := [7]int{23, 4, 25, 64, 1, 242, 21}
//	fmt.Println(array1)
//	x := len(array1) / 2
//	array2 := [3]int{array1[x-1], array1[x], array1[x+1]}
//	fmt.Println(array2)
