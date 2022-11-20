package main

import "fmt"

/*
Given an array of ints, return true if the array contains no 1's and no 3's.

lucky13([0, 2, 4]) → true
lucky13([1, 2, 3]) → false
lucky13([1, 2, 4]) → true
*/

func sliceHasNoOnesAndThrees(givenSlice []int) bool {
	isOne, isThree := false, false
	for i, gs := range givenSlice {
		switch gs {
		case 1:
			isOne = true
		case 3:
			isThree = true
		}
		if i > 1 && isOne == true && isThree == true {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(sliceHasNoOnesAndThrees([]int{2, 2, 2, 3, 4, 2, 1})) //false
}

//func noOneOrThreeInSlice(givenSlice []int) {
//
//	unluckyElement1 := 1
//	unluckyElement2 := 3
//	var luckMeasurer bool
//
//	for _, c := range givenSlice {
//
//		if c != unluckyElement1 && c != unluckyElement2 {
//			luckMeasurer = true
//		}
//		if c == unluckyElement1 || c == unluckyElement2 {
//			luckMeasurer = false
//			break
//		}
//	}
//	fmt.Println(luckMeasurer)
//}
