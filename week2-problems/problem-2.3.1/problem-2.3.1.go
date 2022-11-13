package main

import "fmt"

/*
Given an array of ints, return true if 6 appears as either the first or last element in the array. The array will be length 1 or more.
firstLast6([1, 2, 6]) → true
firstLast6([6, 1, 2, 3]) → true
firstLast6([13, 6, 1, 2, 3]) → false
*/

const length = 6

func firstLast6(givenArray [length]int) bool {
	lastIndex := len(givenArray) - 1
	if givenArray[0] == 6 || givenArray[lastIndex] == 6 {
		return true
	}
	return false
}

func main() {
	fmt.Println(firstLast6([length]int{2, 3, 4, 22, 14, 6})) //true
}

//package main
//
//import "fmt"
//
//func return1(x int, y int) bool {
//	z := false
//	array1 := [6]int{x, 6, 5, 14, 12, y}
//	if array1[0] == 6 || array1[5] == 6 {
//		z = true
//	}
//	return z
//}
//
//func main() {
//	fmt.Println(return1(6, 12))
//}
