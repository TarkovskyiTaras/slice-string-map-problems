package main

import (
	"fmt"
	"log"
)

/*
Given 2 int slices, a and b, both of an odd length, return a new slice length 2 containing their middle elements.

middleWay([1, 2, 3], [4, 5, 6]) → [2, 5]
middleWay([7, 7, 7], [3, 8, 0]) → [7, 8]
middleWay([5, 2, 9], [1, 4, 5]) → [2, 4]
*/

func oddLengthOK(givenSliceOne, givenSliceTwo []int) bool {
	return len(givenSliceOne)%2 != 0 && len(givenSliceTwo)%2 != 0
}

func middleWay(givenSliceOne, givenSliceTwo []int) []int {
	if oddLengthOK(givenSliceOne, givenSliceTwo) == false {
		log.Fatal("the length of the slice(s) doesn't meet the requirement")
	}

	middlePointSliceOne := len(givenSliceOne) / 2
	middlePointSliceTwo := len(givenSliceTwo) / 2
	resultSlice := make([]int, 2)[:2]

	resultSlice[0], resultSlice[1] = givenSliceOne[middlePointSliceOne], givenSliceTwo[middlePointSliceTwo]
	return resultSlice
}

func main() {
	fmt.Println(middleWay([]int{1, 2, 3, 4, 5}, []int{5, 6, 7, 8, 9}))
}

//package main
//
//import "fmt"
//
//func middleWay(array1 [3]int, array2 [3]int) [2]int {
//	array3 := [2]int{array1[1], array2[1]}
//	return array3
//}
//
//func main() {
//	a := [3]int{22, 42, 53}
//	b := [3]int{422, 532, 1024}
//	fmt.Println(middleWay(a, b))
//}
