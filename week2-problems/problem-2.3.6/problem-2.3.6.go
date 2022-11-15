package main

import "fmt"

/*
Given an int slice of any length, if there is a 2 in the slice immediately followed by a 3, set the 3 element to 0. Return the changed slice.

fix23([1, 2, 3]) → [1, 2, 0]
fix23([2, 3, 5]) → [2, 0, 5]
fix23([1, 2, 1]) → [1, 2, 1]
*/

func changeThreeFollowingTwo(givenSlice []int) []int {
	lastIndex := len(givenSlice) - 1
	for i := 1; i <= lastIndex; i++ {
		if givenSlice[i] == 3 && givenSlice[i-1] == 2 {
			givenSlice[i] = 0
		}
	}
	return givenSlice
}

func main() {
	fmt.Println(changeThreeFollowingTwo([]int{2, 2, 3, 4, 2, 3, 3, 123, -1, 2, 3}))
}

//func main() {
//	array1 := [3]int{2, 3, 242}
//	fmt.Println(array1)
//	for i := 0; i < 2; i++ {
//		if array1[i] == 2 || array1[i+1] == 3 {
//			array1[i+1] = 0
//		}
//	}
//	fmt.Println(array1)
//}
