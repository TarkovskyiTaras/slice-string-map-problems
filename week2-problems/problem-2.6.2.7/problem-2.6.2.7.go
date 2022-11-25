package main

import "fmt"

/*
Given a slice of ints, return true if the slice contains either 3 even or 3 odd values all next to each other.

modThree([2, 1, 3, 5]) → true
modThree([2, 1, 2, 5]) → false
modThree([2, 4, 2, 5]) → true
*/

func modThree(givenSlice []int) bool {
	lastIndex := len(givenSlice) - 1
	isEven, isOdd := false, false
	for i := 0; i <= lastIndex-2; i++ {
		for _, gs := range givenSlice[i : i+3] {
			if gs%2 == 0 {
				isEven = true
			} else {
				isOdd = true
			}
		}
		if isEven == true && isOdd == true {
			continue
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(modThree([]int{1, 3, 4, 4, 5, 6, 9}))
}
