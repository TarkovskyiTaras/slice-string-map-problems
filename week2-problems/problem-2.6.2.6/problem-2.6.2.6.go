package main

import (
	"fmt"
	"log"
)

/*
Given slices nums1 and nums2 of the same length, for every element
in nums1, consider the corresponding element in nums2 (at the same index).
Return the count of the number of times that the two elements differ by 2 or less, but are not equal.

matchUp([1, 2, 3], [2, 3, 10]) → 2
matchUp([1, 2, 3], [2, 3, 5]) → 3
matchUp([1, 2, 3], [2, 3, 3]) → 2
*/

func matchUp(givenSliceOne, givenSliceTwo []int) int {
	if len(givenSliceOne) != len(givenSliceTwo) {
		log.Fatal("slices in arguments are not of same length")
	}
	matchCount := 0
	for i := range givenSliceOne {
		elDiff := givenSliceOne[i] - givenSliceTwo[i]
		if elDiff < 0 {
			elDiff = elDiff * (-1)
		}
		if elDiff <= 2 && elDiff != 0 {
			matchCount++
		}
	}
	return matchCount
}

func main() {
	fmt.Println(matchUp([]int{-1, -4, -6}, []int{-2, -6, -8})) // 3
}
