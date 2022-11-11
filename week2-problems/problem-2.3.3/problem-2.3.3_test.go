package main

import "testing"

type testData struct {
	givenSlice []int
	want       []int
}

func twoSlicesEqualChecker(sliceOne []int, sliceTwo []int) bool {

	if len(sliceOne) != len(sliceTwo) {
		return false
	}

	for i, _ := range sliceOne {
		if sliceOne[i] != sliceTwo[i] {
			return false
		}
	}
	return true
}

func TestReverseSlice(t *testing.T) {

	tests := []testData{
		{givenSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			want: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{givenSlice: []int{-24, 5, 1, 0, 123, -6},
			want: []int{-6, 123, 0, 1, 5, -24}},
	}

	for _, test := range tests {
		got := reverseSlice(test.givenSlice)

		if twoSlicesEqualChecker(got, test.want) != true {
			t.Errorf("revereseSlice(\"%#v\") = \"%v\", want: \"%v\"", test.givenSlice, got, test.want)
		}
	}
}
