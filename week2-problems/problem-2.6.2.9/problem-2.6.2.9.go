package main

import (
	"fmt"
)

/*
Return an array that is "left shifted" by one -- so {6, 2, 5, 3} returns {2, 5, 3, 6}.
You may modify and return the given array, or return a new array.

shiftLeft([6, 2, 5, 3]) → [2, 5, 3, 6]
shiftLeft([1, 2]) → [2, 1]
shiftLeft([1]) → [1]
*/

func reverse(slices ...[]int) [][]int {
	for _, sl := range slices {
		for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
			sl[i], sl[j] = sl[j], sl[i]
		}
	}
	return slices
}

func shiftLeft(givenSlice []int, shiftPlaces int) []int {
	reverse(givenSlice[:shiftPlaces], givenSlice[shiftPlaces:], givenSlice)
	return givenSlice
}

func main() {
	fmt.Println(shiftLeft([]int{6, 7, 1, 2, 3, 4, 5}, 4))
}
