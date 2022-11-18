package main

import (
	"testing"
)

type testData struct {
	givenSlice []int
	middleSpan int
	wantSlice  []int
	wantBool   bool
}

func twoSliceEqualChecker(sliceOne []int, sliceTwo []int) bool {
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

func TestFuncDataOK(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{1, 3, 4},
			middleSpan: 3,
			wantBool:   true},
		{givenSlice: []int{1, 2, 3, 4},
			middleSpan: 4,
			wantBool:   false},
		{givenSlice: []int{1, 2, 3, 4, 5},
			middleSpan: 4,
			wantBool:   false},
		{givenSlice: []int{1, 2, 3, 4},
			middleSpan: 3,
			wantBool:   false},
		{givenSlice: []int{1, 2},
			middleSpan: 3,
			wantBool:   false},
		{givenSlice: []int{1, 2, 3, 4, 5},
			middleSpan: 7,
			wantBool:   false},
	}

	for _, test := range tests {
		got := oddLengthOK(test.givenSlice, test.middleSpan) &&
			midSpanNotMoreThanSliceOK(test.givenSlice, test.middleSpan) &&
			sliceLenNotLessThanThreeOK(test.givenSlice, test.middleSpan)

		if got != test.wantBool {
			t.Errorf("oddLengthOK(%#v, %d) = %t, want: %t", test.givenSlice, test.middleSpan, got, test.wantBool)
		}
	}
}

func TestMiddleSlice(t *testing.T) {
	tests := []testData{
		{givenSlice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			middleSpan: 5,
			wantSlice:  []int{3, 4, 5, 6, 7},
		},
		{givenSlice: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			middleSpan: 7,
			wantSlice:  []int{2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test := range tests {
		got := middleSlice(test.givenSlice, test.middleSpan)
		if !twoSliceEqualChecker(got, test.wantSlice) {
			t.Errorf("middleSlice(%#v, %d) = %v, want: %v", test.givenSlice, test.middleSpan, got, test.wantSlice)
		}
	}

}
