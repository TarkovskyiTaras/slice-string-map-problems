package main

import "fmt"

/*
Return true if the group of N numbers at the start and end of the slice are the same.
For example, with {5, 6, 45, 99, 13, 5, 6}, the ends are the same for n=0 and n=2, and false for n=1 and n=3.

sameEnds([5, 6, 45, 99, 13, 5, 6], 1) → false
sameEnds([5, 6, 45, 99, 13, 5, 6], 2) → true
sameEnds([5, 6, 45, 99, 13, 5, 6], 3) → false
*/

func twoSliceEqual(sliceOne []int, sliceTwo []int) bool {
	if len(sliceOne) != len(sliceTwo) {
		return false
	}
	for i := range sliceOne {
		if sliceOne[i] != sliceTwo[i] {
			return false
		}
	}
	return true
}

func sameEnds(givenSlice []int, numbOfEl int) bool {
	leftGroup := givenSlice[:numbOfEl]
	rightGroup := givenSlice[len(givenSlice)-numbOfEl:]
	if twoSliceEqual(leftGroup, rightGroup) {
		return true
	}
	return false
}

func main() {
	fmt.Println(sameEnds([]int{12, 4, 5, 22, 4, 778, 12, 4}, 2)) // true
}
