package main

import "fmt"

/*
We'll say that a 1 immediately followed by a 3 in a slice is an "unlucky" 1. Return true if the given slice contains an unlucky 1 in the first 2 or last 2 positions in the slice.

unlucky1([1, 3, 4, 5]) → true
unlucky1([2, 1, 3, 4, 5]) → true
unlucky1([1, 1, 1]) → false
*/

func unluckyOne(givenSlice []int) bool {
	givenSliceLength := len(givenSlice)
	lastIndex := givenSliceLength - 1

	switch givenSliceLength {
	case 0:
		return false
	case 1:
		return false
	case 2:
		if givenSlice[0] == 1 && givenSlice[1] == 3 {
			return true
		} else {
			return false
		}
	}

	for i := 0; i < 2; i++ {
		if givenSlice[i] == 1 && givenSlice[i+1] == 3 || givenSlice[lastIndex] == 3 && givenSlice[lastIndex-1] == 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(unluckyOne([]int{1, 3, 5, 6, 8})) //true
}

//func main() {
//	status := false
//	array1 := [10]int{3, 1, 0, 0, 1, 3, 0, 0, 1, 3}
//
//	for i := 0; i <= 1; i++ {
//		if array1[i] == 1 && array1[i+1] == 3 {
//			status = true
//		}
//	}
//	for i := len(array1) - 4; i < len(array1)-1; i++ {
//		if array1[i] == 1 && array1[i+1] == 3 {
//			status = true
//		}
//	}
//	fmt.Println(array1)
//	fmt.Println(status)
//}
