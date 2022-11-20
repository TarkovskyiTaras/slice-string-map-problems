package main

import "fmt"

/*
Return the number of even ints in the given slice. Note: the % "mod" operator computes the remainder, e.g. 5 % 2 is 1.

countEvens([2, 1, 2, 3, 4]) → 3
countEvens([2, 2, 0]) → 3
countEvens([1, 3, 5]) → 0
*/

func countEvensInSlice(givenSlice []int) int {
	evenElementsCount := 0
	for i := range givenSlice {
		if givenSlice[i]%2 == 0 {
			evenElementsCount++
		}
	}
	return evenElementsCount
}

func main() {
	fmt.Println(countEvensInSlice([]int{1, 2, 3, 2, 3, 2, 0}))
}

//func main() {
//	givenArray := [10]int{1, 43, 3, 4, 5, 6, 7, 8, 9, 101}
//	countEven := 0
//	for i, _ := range givenArray {
//		if givenArray[i]%2 == 0 {
//			countEven++
//		}
//	}
//	fmt.Println(countEven)
//}
