package main

import (
	"errors"
	"fmt"
	"log"
)

/*
Given a non-empty slice of ints, return a new slice containing the elements from the original slice that come after
the second last 4 in the original slice. The original slice will contain at least two 4s.

post4([2, 4, 1, 2]) → [1, 2]
post4([4, 1, 4, 2]) → [2]
post4([4, 4, 1, 2, 3]) → [1, 2, 3]
*/

func indexAtSecondFour(givenSlice []int) (int, error) {
	if err := errors.New("invalid slice length"); len(givenSlice) < 2 {
		log.Fatal(err.Error())
	}

	fourCount := 0
	lastIndex := len(givenSlice) - 1
	for j := lastIndex; j >= 0; j-- {
		if givenSlice[j] == 4 {
			fourCount++
		}
		if fourCount == 2 {
			return j, nil
		}
	}
	return -1, errors.New("at least two 4s requirement not met")
}

func post4Slice(givensSlice []int) []int {
	index, err := indexAtSecondFour(givensSlice)
	if err != nil {
		log.Fatal(err)
	}
	return givensSlice[index+1:]
}

func main() {
	fmt.Println(post4Slice([]int{1, 2, 3, 4, 5, 6, 4, 1, 2, 3})) //[5  6  4 1 2 3]
}
