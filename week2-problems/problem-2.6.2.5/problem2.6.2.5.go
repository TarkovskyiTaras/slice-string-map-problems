package main

import "fmt"

/*
Given a slice of ints, return true if it contains no 1's and no 4's.

no14([1, 2, 3]) → true
no14([1, 2, 3, 4]) → false
no14([2, 3, 4]) → true
*/

func noOneOrFourInSlice(givenSlice []int) bool {
	if len(givenSlice) < 2 {
		return true
	}
	var isOne, isFour bool
	for _, gs := range givenSlice {
		switch gs {
		case 1:
			isOne = true
		case 4:
			isFour = true
		}
		if isOne && isFour {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(noOneOrFourInSlice([]int{1, 2, 3, 4, 5})) //false
}

//func noOneOrFourInSlice(givenSlice []int) {
//	unluckyNumb1 := 1
//	unluckyNumb2 := 4
//	luckMeasurer := true
//	for _, c := range givenSlice {
//		if c == unluckyNumb1 || c == unluckyNumb2 {
//			luckMeasurer = false
//			break
//		}
//
//	}
//	fmt.Println(luckMeasurer)
//}
