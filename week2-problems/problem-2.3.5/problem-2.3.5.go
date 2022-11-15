package main

import "fmt"

/*
Given an int slice of any length, return true if it does not contain a 2 or 3.

no23([4, 5]) → true
no23([4, 2]) → false
no23([3, 5]) → false
*/

func noTwoOrThreeInSlice(givenSlice []int) bool {
	for _, x := range givenSlice {
		if x == 2 || x == 3 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(noTwoOrThreeInSlice([]int{1, 4, 6, 8})) //true
}

//func no23(x int, y int) bool {
//	boolZ := true
//	array1 := [2]int{x, y}
//	for _, z := range array1 {
//		if z == 2 || z == 3 {
//			boolZ = false
//		}
//	}
//
//	return boolZ
//}
//
//func main() {
//	fmt.Println(no23(2, 23))
//}
