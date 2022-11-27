package main

import "testing"

type testData struct {
	givenSliceOneFuncOddLengthOK, givenSliceTwoFuncOddLengthOk []int
	wantFuncOddLength                                          bool
	givenSliceOneFuncMiddleWay, givenSliceTwoFuncMiddleWay     []int
	wantFuncMiddleWay                                          []int
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

func TestOddLengthOK(t *testing.T) {
	tests := []testData{
		{givenSliceOneFuncOddLengthOK: []int{1, 2, 3},
			givenSliceTwoFuncOddLengthOk: []int{4, 5, 6},
			wantFuncOddLength:            true,
		},
		{givenSliceOneFuncOddLengthOK: []int{1, 2, 4},
			givenSliceTwoFuncOddLengthOk: []int{7, 8},
			wantFuncOddLength:            false,
		},
		{givenSliceOneFuncOddLengthOK: []int{7, -12},
			givenSliceTwoFuncOddLengthOk: []int{77, 6, 4},
			wantFuncOddLength:            false,
		},
		{givenSliceOneFuncOddLengthOK: []int{1, 6},
			givenSliceTwoFuncOddLengthOk: []int{2, 2},
			wantFuncOddLength:            false,
		}}

	for _, test := range tests {
		got := oddLengthOK(test.givenSliceOneFuncOddLengthOK, test.givenSliceTwoFuncOddLengthOk)
		if got != test.wantFuncOddLength {
			t.Errorf("oddLengthOK(\"%#v\", \"%#v\") = \"%t\", want: \"%t\"",
				test.givenSliceOneFuncOddLengthOK, test.givenSliceTwoFuncOddLengthOk, got, test.wantFuncOddLength)
		}
	}
}

func TestMiddleWay(t *testing.T) {
	tests := []testData{
		{givenSliceOneFuncMiddleWay: []int{1, 2, 3},
			givenSliceTwoFuncMiddleWay: []int{4, 5, 6, 7, 8},
			wantFuncMiddleWay:          []int{2, 6},
		},
		{givenSliceOneFuncMiddleWay: []int{1, -4, 3, 5, 6},
			givenSliceTwoFuncMiddleWay: []int{4, 5, -6, 7, -8},
			wantFuncMiddleWay:          []int{3, -6},
		},
	}
	for _, test := range tests {
		got := middleWay(test.givenSliceOneFuncMiddleWay, test.givenSliceTwoFuncMiddleWay)

		if twoSliceEqualChecker(got, test.wantFuncMiddleWay) != true {
			t.Errorf("middleWay(\"%#v\", \"%#v\") = \"%#v\", want: \"%#v\"",
				test.givenSliceOneFuncMiddleWay, test.givenSliceTwoFuncMiddleWay, got, test.wantFuncMiddleWay)
		}
	}
}
