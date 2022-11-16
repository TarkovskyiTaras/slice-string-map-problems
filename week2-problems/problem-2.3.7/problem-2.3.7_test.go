package main

import (
	"testing"
)

type testData struct {
	givenSlice   []int
	want         []int
	passedDataOK bool
}

func twoSliceEqualChecker(sliceOne, sliceTwo []int) bool {
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

func TestPassedDataOK(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{1},
			passedDataOK: false},
		{givenSlice: []int{1, 2, 3},
			passedDataOK: false},
		{givenSlice: []int{1, 2, 3, 4, 5, 6},
			passedDataOK: true},
	}
	for _, test := range tests {
		got := evenLengthOK(test.givenSlice) && lengthTwoOrMoreOK(test.givenSlice)
		if got != test.passedDataOK {
			t.Errorf("data passed to the function doesn't meet the requirements")
		}
	}
}

func TestMiddleMake(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{2, 3},
			want: []int{2, 3}},
		{givenSlice: []int{22, 12, -5, 6, 15, 4},
			want: []int{-5, 6}},
	}
	for _, test := range tests {
		got := makeMiddle(test.givenSlice)
		if !twoSliceEqualChecker(got, test.want) {
			t.Errorf("makeMiddle(%#v) = \"%v\", want: \"%v\"", test.givenSlice, got, test.want)
		}
	}
}
