package main

import "fmt"

/*
Given an array of ints length 3 (slice of ints), return a new array with the elements in reverse order, so {1, 2, 3} becomes {3, 2, 1}.

reverse3([1, 2, 3]) → [3, 2, 1]
reverse3([5, 11, 9]) → [9, 11, 5]
reverse3([7, 0, 0]) → [0, 0, 7]
*/

func reverseSlice(givenSlice []int) []int {
	lastIndex := len(givenSlice) - 1
	for i, j := 0, lastIndex; j > i; i, j = i+1, j-1 {
		givenSlice[i], givenSlice[j] = givenSlice[j], givenSlice[i]
	}
	return givenSlice
}

func main() {
	fmt.Println(reverseSlice([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})) //[9 8 7 6 5 4 3 2 1 0]
}

//func reverse3(x int, y int, z int) [3]int {
//	array1 := [3]int{x, y, z}
//	fmt.Println(array1)
//	arrayReverse := [3]int{z, y, x}
//
//	return arrayReverse
//}
//
//func main() {
//	fmt.Println(reverse3(21, 256, 424))
//}
